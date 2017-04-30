package gophoenix

import "sync"

type messageRouter struct {
	mapLock sync.RWMutex
	tr map[string]*topicReceiver
	sub chan ChannelReceiver
}

type topicReceiver struct {
	cr ChannelReceiver
	rr *replyRouter
}

func NewMessageRouter() *messageRouter {
	return &messageRouter{
		tr: make(map[string]*topicReceiver),
		sub: make(chan ChannelReceiver),
	}
}

func (mr messageRouter) notifyMessage(msg *Message) {
	mr.mapLock.RLock()
	tr, ok := mr.tr[msg.Topic]
	mr.mapLock.Unlock()
	if !ok {
		return
	}

	switch msg.Event {
	case replyEvent:
		tr.rr.routeReply(msg)
	case joinEvent:
		tr.cr.OnJoin(msg.Payload)
	case errorEvent:
		tr.cr.OnJoinError(msg.Payload)
	case closeEvent:
		tr.cr.OnChannelClose(msg.Payload)
	default:
		tr.cr.OnMessage(msg.Event, msg.Payload)
	}
}

func (mr messageRouter) subscribe(topic string, cr ChannelReceiver, rr *replyRouter) {
	mr.mapLock.Lock()
	defer mr.mapLock.Unlock()
	mr.tr[topic] = &topicReceiver{cr: cr, rr: rr}
}
