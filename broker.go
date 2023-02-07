package pubsub

import (
	"log"
	"sync"
)

type Subscribers map[string]*Subscriber // id to subscriber map
type TopicHandler func(s *Subscriber, m *Message)

type Broker struct {
	subscribers   Subscribers // for braodcasting messages to all subscriber independent of topics
	topics        map[string]Subscribers
	mut           sync.RWMutex
	name          string // name has not any functional impact, we can skip it
	topicHandlers map[string]TopicHandler
}

func NewBroker(name string) *Broker {
	return &Broker{
		subscribers:   make(Subscribers),
		topics:        make(map[string]Subscribers),
		name:          name,
		topicHandlers: make(map[string]TopicHandler),
	}
}

func (b *Broker) AddSubscriber() *Subscriber {
	b.mut.Lock()
	defer b.mut.Unlock()
	subscriber := CreateNewSubscriber()

	for b.subscribers[subscriber.Id] != nil {
		subscriber = CreateNewSubscriber()
	}

	b.subscribers[subscriber.Id] = subscriber

	return subscriber

}

func (b *Broker) Subscribe(s *Subscriber, topic string) bool {

	b.mut.Lock()
	defer b.mut.Unlock()

	subscriberId := s.Id

	if b.topics[topic] == nil {
		b.topics[topic] = make(Subscribers)
	}

	s.AddTopic(topic)
	b.topics[topic][subscriberId] = s

	log.Printf("Subscribed to a topic %v\n", topic)

	return true
}

func (b *Broker) Unsubscribe(s *Subscriber, topic string) {
	b.mut.Lock()
	defer b.mut.Lock()

	delete(b.topics[topic], s.Id)
	s.RemoveTopic(topic)

	log.Printf("Unsubscribed to a topic %v\n", topic)

}

func (b *Broker) Publish(topic string, msg string) {
	// 1. get subscribers by topic
	b.mut.RLock()
	topSubs := b.topics[topic]
	b.mut.RUnlock()

	// 2. loop to topic subscribers and signal to each subscriber
	for _, sub := range topSubs {
		go sub.Signal(NewMessage(topic, msg))
	}
}

func (b *Broker) Broadcast(msg string) {
	m := NewMessage("", msg)
	for _, sub := range b.subscribers {
		go sub.Signal(m)
	}
}

func (b *Broker) RemoveSubscriber(s *Subscriber) {
	b.mut.Lock()
	defer b.mut.Unlock()

	for topic := range s.Topics {
		b.Unsubscribe(s, topic)
	}

	s.Destruct()
	delete(b.subscribers, s.Id)

}

func (b *Broker) SetHandler(topic string, handler TopicHandler) {
	b.mut.Lock()
	defer b.mut.Unlock()

	b.topicHandlers[topic] = handler

}

func (b *Broker) GetHandlerByTopic(topic string) TopicHandler {
	return b.topicHandlers[topic]
}
