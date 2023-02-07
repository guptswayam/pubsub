package pubsub

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Subscriber struct {
	Id          string
	MessageChan chan *Message
	Topics      map[string]bool
	Active      bool
	mutex       sync.RWMutex
}

func CreateNewSubscriber() *Subscriber {
	rand.Seed(time.Now().UnixMicro())

	id := rand.Int()

	return &Subscriber{
		Id:          fmt.Sprint(id),
		MessageChan: make(chan *Message),
		Topics:      make(map[string]bool),
		Active:      true,
	}

}

func (s *Subscriber) RemoveTopic(topic string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	delete(s.Topics, topic)
}

func (s *Subscriber) AddTopic(topic string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Topics[topic] = true
}

func (s *Subscriber) Signal(msg *Message) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.Active {
		s.MessageChan <- msg
	}
}

func (s *Subscriber) Destruct() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.Active = false
	close(s.MessageChan)
}

func (s *Subscriber) Listen(handler TopicHandler) {
	for {
		msg := <-s.MessageChan
		handler(s, msg)
	}
}
