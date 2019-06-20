package pubsubtest

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

func ExampleMyStruct_realExample() {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		log.Fatal(err)
		return
	}
	topic, err := pubsubClient.CreateTopic(context.Background(), "topic-name")
	ms := NewMyStruct(topicWrapper{topic})
	ms.DoSomething(ctx)
}

type stubReadyGetter struct {
	ServerID string
	Err      error
	ReadyC   <-chan struct{}
}

func (frg stubReadyGetter) Get(context.Context) (serverID string, err error) {
	<-frg.ReadyC
	return frg.ServerID, frg.Err
}

func (frg stubReadyGetter) Ready() <-chan struct{} {
	return frg.ReadyC
}

type stubPublisher struct {
	Result readyGetter
}

func (fp stubPublisher) Publish(_ context.Context, msg *pubsub.Message) readyGetter {
	return fp.Result
}

func ExampleMyStruct_testWithStub() {
	ready := make(chan struct{})
	m := NewMyStruct(stubPublisher{
		Result: stubReadyGetter{
			ServerID: "myServerID",
			Err:      nil,
			ReadyC:   ready,
		},
	})
	go func() {
		time.Sleep(10 * time.Millisecond)
		close(ready)
	}()
	m.DoSomething(context.Background())
	// Output:
	// tick
	// tick
	// serverID = myServerID
}
