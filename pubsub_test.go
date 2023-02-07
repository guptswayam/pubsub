package pubsub_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/guptswayam/pubsub"
)

func Test_PubSub(t *testing.T) {

	broker := pubsub.NewBroker("test-broker")
	topic1 := "test"

	broker.SetHandler(topic1, func(s *pubsub.Subscriber, msg *pubsub.Message) {
		log.Println("subscriber with id '"+s.Id+"' recieved message:", msg)
	})

	s1 := broker.AddSubscriber()
	s2 := broker.AddSubscriber()

	broker.Subscribe(s1, topic1)
	broker.Subscribe(s2, topic1)

	go s1.Listen(broker.GetHandlerByTopic(topic1))
	go s2.Listen(broker.GetHandlerByTopic(topic1))

	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		broker.Publish(topic1, "Hi there! - "+fmt.Sprint(i+1))
	}

	time.Sleep(10 * time.Millisecond)
	t.Log("Success")
}
