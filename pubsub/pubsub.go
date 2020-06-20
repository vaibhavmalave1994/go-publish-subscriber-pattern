package pubsub

import "time"

type Bus struct {
	Subscribers map[string][]*Subscriber // map[topic][0...n]*subcribers
}

type Subscriber struct {
	Channel <-chan *Message //channel from where subscriber will read messages
	Closed  bool            // to check whether channel is closed or not
	ID      string          //add id to subscriber to identify which subscriber it is
}

type Publisher struct{}

type Message struct {
	payload []byte
}

func NewPubSub() *Bus {
	return &Bus{
		make(map[string][]*Subscriber),
	}
}

func (b *Bus) Subscribe(topic string) (<-chan *Message, error) {
	//create new output channel for subscriber
	messageChannel := make(<-chan *Message)

	//create new subscriber
	subscriber := &Subscriber{
		Channel: messageChannel,
		Closed:  false,
		ID:      string(time.Now().UnixNano()),
	}

	//assign subscriber to given topic
	if _, ok := b.Subscribers[topic]; !ok {
		b.Subscribers[topic] = make([]*Subscriber, 0)
	}
	b.Subscribers[topic] = append(b.Subscribers[topic], subscriber)

	//return output channel

	return messageChannel, nil

}
