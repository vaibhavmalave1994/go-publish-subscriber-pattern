package pubsub

import (
	"fmt"
	"testing"
)

func TestNewPubSub(t *testing.T) {
	actualPubSub := NewPubSub()

	if actualPubSub == nil {
		t.Error("Expcted pointer to Bus after calling NewPubSub instead received nil")
	}
}

func TestSubscribe(t *testing.T) {
	pubSub := NewPubSub()

	subscriberOne, err := pubSub.Subscribe("example.Topic")

	subscriberTow, err := pubSub.Subscribe("example.Topic")

	if subscriberOne == nil {
		t.Error("Expcted output channel after calling Subscribe instead received nil")
	}

	if err != nil {
		t.Errorf("Expcted no error after calling Subscribe instead received error %v", err)
	}

	message := &Message{
		[]byte(string("hello world")),
	}
	fmt.Println("message")
	sendMessageErr := pubSub.Publish("example.Topic", message)

	if sendMessageErr != nil {
		t.Errorf("Expcted no error after calling SendMessage instead received error %v", err)
	}
	for {
		go func(outputChan <-chan *Message) {
			for message := range outputChan {
				fmt.Println("message 1", string(message.payload))
			}
		}(subscriberOne)

		go func(outputChan <-chan *Message) {
			for message := range outputChan {
				fmt.Println("message 2", string(message.payload))
			}
		}(subscriberTow)
	}

}
