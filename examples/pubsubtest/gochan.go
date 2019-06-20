package pubsubtest

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

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

// SyncPublisher can publish msg synchronously.
type SyncPublisher interface {
	PublishSync(ctx context.Context, msg *pubsub.Message) (serverID string, err error)
}

// DoSomethingUsingSyncPublisher does something using synchronous publisher.
func DoSomethingUsingSyncPublisher(ctx context.Context, sp SyncPublisher) {
	// Build message.
	msg := &pubsub.Message{}
	publishDone := make(chan struct{})
	allDone := make(chan struct{})
	go func() {
		defer close(allDone)
		serverID, err := sp.PublishSync(ctx, msg)
		close(publishDone)
		// Do something with serverID and err.
		if err != nil {
			log.Fatal(err)
		}
		// This may take a while (4ms). But tick is not going to be
		// printed anymore.
		time.Sleep(4 * time.Millisecond)
		fmt.Printf("serverID = %s\n", serverID)
	}()
	ticker := time.NewTicker(4 * time.Millisecond)
	for {
		select {
		case <-publishDone:
			break
		case <-ticker.C:
			fmt.Println("tick")
			continue
		}
		break
	}
	ticker.Stop()
	<-allDone
}
