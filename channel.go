package gophoenix

import (
	"encoding/json"
	"fmt"
)

type Channel struct {
	topic string
	t Transport
	rc refCounter
	rr *replyRouter
}

type refCounter interface {
	nextRef() int64
}

func (ch *Channel) Join(payload interface{}) error {
	ref := ch.rc.nextRef()
	return ch.sendMessage(ref, joinEvent, payload)
}

func (ch *Channel) Leave(payload interface{}) error {
	ref := ch.rc.nextRef()
	return ch.sendMessage(ref, leaveEvent, payload)
}

func (ch *Channel) Push(event event, payload interface{}, replyHandler func(payload interface{})) error {
	ref := ch.rc.nextRef()
	ch.rr.subscribe()
	return ch.sendMessage(ref, event, payload)
}

func (ch *Channel) sendMessage(ref int64, event event, payload interface{}) error {
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
