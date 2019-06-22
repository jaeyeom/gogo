# 사례 연구: PubSub 테스팅

인터페이스를 이용하면 간단한 경우에는 쉽게 테스트 할 수 있으나 어떤 경우에는
까다로운 경우가 있다. Google의 PubSub API에서
[`Publish()`](https://godoc.org/cloud.google.com/go/pubsub#Topic.Publish) 함수를
찾아 보자. 아래와 같이 `Publish()` 결과로 `*PublishResult`를 반환한다.

```go
func (t *Topic) Publish(ctx context.Context, msg *Message) *PublishResult
```

`Publish()` 호출은 비동기 호출이다. 반환되는 `PublishResult`에 `Get()` 메서드를
호출하여 동기화 호출을 할 수 있고 `Ready()`에서 채널을 넘겨 받아서 비동기 처리를
해 줄 수도 있다.

```go
// Synchronous call
serverID, err := topic.Publish(ctx, msg).Get(ctx)

// Asynchronous call
result := topic.Publish(ctx, msg)
for {
  select {
  case <-result.Ready():
    serverID, error := result.Get(ctx)
    // Do something.
  case ...:
    // ...
  default:
  }
}
```

일단은 왜 이렇게 API를 복잡하게 만들었는지 잘 이해가 되지 않는다. 그냥 동기화
호출만 제공하더라도 매우 쉽게 비동기 호출이 가능하다. 즉 `Publish()`가 그냥
동기화 호출이라도 상관없다는 얘기다. 그냥 함수 호출 할 때 앞에 `go`만 붙여주면
쉽게 비동기 호출을 할 수 있고, 고루틴이 언제 끝나는지 즉, 언제 `Ready` 되는지는
채널을 이용하면 쉽게 해결된다.

```go
done := make(chan struct{})
go func() {
  defer close(done)
  serverID, err := topic.Publish(ctx, msg).Get(ctx)
  // Do something
}()
```

이제 테스트를 해야 한다. 처음에는 간단해 보인다. 이 `Publish()` 메서드를
`Publisher` 인터페이스에 담아서 정의한 후 가짜 구현을 하든 Mock을 만들든 하면 될
것 같다. 그러나 안타깝게도 `PublishResult`를 쉽게 생성할 수 있는 방법이 없다.
모든 필드가 비노출이며, `Get()`이나 `Ready()`를 호출했을 때 어떤 값을 반환할지를
테스트 상황에서 컨트롤 할 수 없는 구조다.

한 가지 방법은 어차피 동기화 호출만 있어도 비동기 구현이 가능하므로
`SyncPublisher`라고 하는 인터페이스와 구현을 만드는 것이다. 이것은 비교적
간단하게 할 수 있고, `gochan.go`와 `gochan_test.go` 파일에 시연되어 있다. 이렇게
하지 않고 꼭 인터페이스를 이용하여 Ready()와 Get() 메서드를 호출하는 방식으로
구현하고 싶다면, 더 복잡해지지만 `interface.go`와 `interface_test.go` 파일을
참고하자.

이제 왠만한 테스트는 쉽게 할 수 있을 것이다. 두 방법 모두 mock 패키지를 이용하여
테스트 할 수도 있다.

## 우직하게 테스트 하기

먼저 우직하게 테스트하는 방법부터 알아보자.

일단 복잡한 무언가가 있을 때는 가장 바깥부터 시작하는 것이 정석이라는 것이 그냥
저자의 생각이다. 이 상황에서는 `*PublishResult`를 먼저 공략해 보자. 여기서
`PublishResult`가 무슨 메서드를 제공하는 지가 중요한 것이 아니라 우리가 어떤
메서드를 이용하는지가 중요하다. 물론 이 예제에서는 `*PublishResult`가 제공하는
두 메서드 모두 이용하는 경우이므로 큰 차이가 없지만 여기서 정의하는 인터페이스는
제공되는 메서드가 아니라 사용하는 메서드 위주로 구성해야 한다는 것을 명심해야
한다.

```go
type readyGetter interface {
	Get(ctx context.Context) (serverID string, err error)
	Ready() <-chan struct{}
}
```
그러면 이 인터페이스를 구현하는 무언가를 받아서 작업할 수 있게 만들어야 한다.
실제 구현에서는 진짜 `*PublishResult`를 이용하고 테스트에서는 저 두 메서드를
구현하는 다른 구현체를 이용해도 될 것이다. 그렇다면 저 인터페이스를 반환하는
메서드가 필요하다. 안타깝게도 실제 `pubsub.Topic`이 `Publish`를 할 때는
`*PublishResult`를 반환하기 때문에 이 부분에서 어떻게 제어할 수 없다. 따라서
대신에 아래와 같은 인터페이스를 반환하는 인터페이스를 정의할 필요가 생긴다.
왜 이래야 하는가 생각될 수 있지만, 좀 더 생각해 보면 어쩔 수 없이 이렇게 해야
한다는 결론이 나온다. 슬프지만 이렇게 해 보자.

```go
type publisher interface {
	Publish(ctx context.Context, msg *pubsub.Message) readyGetter
}
```

이것은 `Pubsub.Topic`이 제공하는 `Publish` 함수와 비슷하지만, 반환 값의 자료형이
인터페이스를 반환하므로 다르다.

이제 이 시그니처에 맞는 메서드 선언이 필요한데, 여기서 래퍼가 필요하다. 힙합
대통령이 아니라 원래 토픽을 감싸서 저 형태의 메서드를 구현하는 무언가가
필요하다는 것이다. 물론 이런 것을 만들어 줘야 한다는 것은 매우 성가신 일이고
하기 싫은 일이라는 것은 동감한다.

```go
type topicWrapper struct {
	*pubsub.Topic
}

func (t topicWrapper) Publish(ctx context.Context, msg *pubsub.Message) readyGetter {
	result := t.Topic.Publish(ctx, msg)
	if result == nil {
		panic("never happens")
	}
	return result
}
```

별로 재미는 없지만 위와 같이 토픽을 감싸고 `Publish`의 반환값 자료형을 바꿔서
그저 별 의미없이 반환해 보았다. 문서를 읽어보면 `Publish` 메서드는 절대 `nil`
값을 반환하지 않는다고 하므로 그냥 `return t.Topic.Publish(ctx, msg)`를
호출하여도 동일하다. 그저 한 번 더 방어적인 느낌적인 느낌으로 `nil`값을
반환하는지 한 번 더 검사를 해 주는 것일 뿐이다.

이제 구현체는 `publisher` 인터페이스를 구현하는 무언가를 받아서 생성되기만 하면
된다.

테스트에서는 `publisher` 인터페이스를 구현하는 무엇이든 넘겨줄 수 있고, 여기서
반환되는 resultGetter가 무엇을 반환할 것이지 역시 제어가 가능해 진다. 나머지는
소스 코드를 확인하면 이해할 수 있을 것이다.

## API를 단순화하기

구글이 API를 단순하게 제공해 줬으면 더 좋았겠다는 생각을 한다. 물론 비동기
호출을 할 수 있게 제공해 준 것은 고맙지만 테스트 할 때 더 힘들어질 분이라는 것을
만든 사람들은 알랑가 모르겠다.

일단 API를 단순화하면 다음과 같이 만들 수 있다.

```go
// SyncTopic provides an additional method PublishSync() which is a synchronous
// version of Publish().
type SyncTopic struct {
	*pubsub.Topic
}

// PublishSync publishes to the underlying topic synchronously. Google should've
// made Publish() synchronous like this because it's easy to call asynchronously
// using synchronous API and it's easier to test.
func (st SyncTopic) PublishSync(ctx context.Context, msg *pubsub.Message) (serverID string, err error) {
	return st.Topic.Publish(ctx, msg).Get(ctx)
}
```

역시 원래 있던 토픽을 감싸는 것은 같지만 이번에는 `Publish`를 구현하는 것이
아니라 동기화 API를 구현한다. 굳이 복잡하게 할 필요 없이 `Publish` 이후에 `Get`
함수를 호출하자는 것이다. 여기서 `Get` 함수는 동기화 호출이므로 결과값이 반환될
때까지 기다려야 하지만, Go 언어는 말 그대로 `go` 키워드를 이용하여 비동기 호출을
쉽게 할 수 있다. 이 동기화 호출은 바로 우리가 원하는 값을 반환받을 수 있어서
더더욱 편리하다.

이제 이 동기화 호출을 인터페이스로 정의해 보자.

```go
// SyncPublisher can publish msg synchronously.
type SyncPublisher interface {
	PublishSync(ctx context.Context, msg *pubsub.Message) (serverID string, err error)
}

```

이렇게 된다면 매우 간단하다. 평소에 다른 함수 테스트하듯이 구현체에서는 이
인터페이스를 받아서 이용하면 되고 테스트에서는 다른 구현을 넘겨주면 된다.

특별할 것이 없으므로 소스 코드를 참조하면 될 것 같다.
