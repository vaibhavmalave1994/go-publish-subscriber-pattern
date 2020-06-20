package pubsub

import (
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

	outputChan, err := pubSub.Subscribe("example.Topic")

	if outputChan == nil {
		t.Error("Expcted output channel after calling Subscribe instead received nil")
	}

	if err != nil {
		t.Errorf("Expcted no error after calling Subscribe instead received error %v", err)
	}

}
