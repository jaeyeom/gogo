package pubsubtest

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

func ExampleDoSomethingUsingSyncPublisher_realExample() {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		log.Fatal(err)
		return
	}
	topic, err := pubsubClient.CreateTopic(context.Background(), "topic-name")
	st := SyncTopic{topic}
	DoSomethingUsingSyncPublisher(ctx, st)
}

type stubSyncPublisher struct {
	ServerID string
	Err      error
	Delay    time.Duration
}

func (sp stubSyncPublisher) PublishSync(ctx context.Context, msg *pubsub.Message) (serverID string, err error) {
	time.Sleep(sp.Delay)
	return sp.ServerID, sp.Err
}

func ExampleDoSomethingUsingSyncPublisher_testUsingStub() {
	sp := stubSyncPublisher{
		ServerID: "myServerID",
		Err:      nil,
		Delay:    10 * time.Millisecond,
	}
	DoSomethingUsingSyncPublisher(context.Background(), sp)
	// Output:
	// tick
	// tick
	// serverID = myServerID
}
