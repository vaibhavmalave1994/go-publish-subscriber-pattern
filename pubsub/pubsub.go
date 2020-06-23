package pubsub

import (
	"fmt"
	"log"
	"math/rand"
)

type Bus struct {
	Subscribers map[string][]*Subscriber // map[topic][0...n]*subcribers
}

type Subscriber struct {
	Channel chan *Message //channel from where subscriber will read messages
	Closed  bool          // to check whether channel is closed or not
	ID      int           //add id to subscriber to identify which subscriber it is
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
	messageChannel := make(chan *Message)

	//create new subscriber
	subscriber := &Subscriber{
		Channel: messageChannel,
		Closed:  false,
		ID:      rand.Int(),
	}

	//assign subscriber to given topic
	if _, ok := b.Subscribers[topic]; !ok {
		b.Subscribers[topic] = make([]*Subscriber, 0)
	}
	b.Subscribers[topic] = append(b.Subscribers[topic], subscriber)

	//return output channel
	return subscriber.Channel, nil

}

func (b *Bus) Publish(topic string, message *Message) error {
	subscribers, ok := b.Subscribers[topic]
	if !ok || len(subscribers) == 0 {
		log.Println("no topic for subscriber")
		return nil
	}
	log.Printf("number of subscribers %d\n", len(subscribers))
	for i := range subscribers {
		subscriber := subscribers[i]
		log.Printf("number of subscribers %s %d\n", string(message.payload), i)
		if subscriber.Closed {
			log.Println("pub/sub closed.")
			return nil
		}
		fmt.Println("ok4")
		go func(subscriber *Subscriber) {
			select {
			case subscriber.Channel <- message:
				log.Printf("Message has been send to subscriber %d \n", subscriber.ID)
				//unblock
			}
		}(subscriber)

	}
	log.Println("done")
	return nil
}
