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
serverID, err := topic.Publish(ctx, msg).Get()

// Asynchronous call
result := topic.Publish(ctx, msg)
for {
  select {
  case <-result.Ready():
    serverID, error := result.Get()
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
  serverID, err := topic.Publish(ctx, msg).Get()
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
간단하게 할 수 있고, `gochan.go`와 `gochan_test.goc` 파일에 시연되어 있다.
이렇게 하지 않고 꼭 인터페이스를 이용하여 Ready()와 Get() 메서드를 호출하는
방식으로 구현하고 싶다면, 더 복잡해지지만 `interface.go`와 `interface_test.go`
파일을 참고하자.

이제 왠만한 테스트는 쉽게 할 수 있을 것이다. 두 방법 모두 mock 패키지를 이용하여
테스트 할 수도 있다.
