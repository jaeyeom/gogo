package pubsubtest

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

type readyGetter interface {
	Get(ctx context.Context) (serverID string, err error)
	Ready() <-chan struct{}
}

type publisher interface {
	Publish(ctx context.Context, msg *pubsub.Message) readyGetter
}

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

type MyStruct struct {
	topic publisher
}

func NewMyStruct(topic publisher) *MyStruct {
	return &MyStruct{
		topic: topic,
	}
}

// DoSomething actually does something using asynchronous feature of pubsub
// publisher.
func (m MyStruct) DoSomething(ctx context.Context) {
	// Build message.
	msg := &pubsub.Message{}
	result := m.topic.Publish(ctx, msg)
	ticker := time.NewTicker(4 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-result.Ready():
			serverID, err := result.Get(ctx)
			if err != nil {
				log.Fatal(err)
			}
			// This may take a while (4ms). But tick is not going to
			// be printed anymore.
			time.Sleep(4 * time.Millisecond)
			fmt.Printf("serverID = %s\n", serverID)
			return
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}
