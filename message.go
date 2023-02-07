package pubsub

type Message struct {
	Topic string
	Body  string
}

func NewMessage(topic, msg string) *Message {
	return &Message{
		Topic: topic,
		Body:  msg,
	}
}
