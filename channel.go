package gophoenix

import (
	"encoding/json"
	"fmt"
)

type Channel struct {
	topic string
	t Transport
	rc refCounter
}

type refCounter interface {
	nextRef() int64
}

func (ch *Channel) Join(payload interface{}) {
	ch.Push(joinEvent, payload)
}

func (ch *Channel) Leave(payload interface{}) {
	ch.Push(leaveEvent, payload)
}

func (ch *Channel) Push(event event, payload interface{}) error {
	ref := ch.rc.nextRef()
	msg := &Message{
		Topic: ch.topic,
		Event: event,
		Payload: payload,
		Ref: ref,
	}

	data, err := json.Marshal(msg)

	if err != nil {
		return fmt.Errorf("unable to marshal message: %s", err)
	}

	return ch.t.Push(data)
}
