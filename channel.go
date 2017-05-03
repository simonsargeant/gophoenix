package gophoenix

import (
	"encoding/json"
	"fmt"
)

// Channel represents a subscription to a topic. It is returned from the Client after joining a topic.
type Channel struct {
	topic string
	t Transport
	rc refCounter
	rr *replyRouter
	ln leaveNotifier
}

type leaveNotifier func ()

type refCounter interface {
	nextRef() int64
}

// Leave notifies the channel to unsubscribe from messages on the topic.
func (ch *Channel) Leave(payload interface{}) error {
	defer ch.ln()
	ref := ch.rc.nextRef()
	return ch.sendMessage(ref, LeaveEvent, payload)
}

// Push sends a message on the topic.
func (ch *Channel) Push(event Event, payload interface{}, replyHandler func(payload interface{})) error {
	ref := ch.rc.nextRef()
	ch.rr.subscribe(ref, replyHandler)
	return ch.sendMessage(ref, event, payload)
}

// PushNoReply sends a message on the topic but does not provide a callback to receive replies.
func (ch *Channel) PushNoReply(event Event, payload interface{}) error {
	ref := ch.rc.nextRef()
	return ch.sendMessage(ref, event, payload)
}

func (ch *Channel) join(payload interface{}) error {
	ref := ch.rc.nextRef()
	return ch.sendMessage(ref, JoinEvent, payload)
}

func (ch *Channel) sendMessage(ref int64, event Event, payload interface{}) error {
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
