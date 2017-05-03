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

func newMessageRouter() *messageRouter {
	return &messageRouter{
		tr: make(map[string]*topicReceiver),
		sub: make(chan ChannelReceiver),
	}
}

func (mr *messageRouter) NotifyMessage(msg *Message) {
	mr.mapLock.RLock()
	tr, ok := mr.tr[msg.Topic]
	mr.mapLock.Unlock()
	if !ok {
		return
	}

	switch msg.Event {
	case ReplyEvent:
		tr.rr.routeReply(msg)
	case JoinEvent:
		tr.cr.OnJoin(msg.Payload)
	case ErrorEvent:
		tr.cr.OnJoinError(msg.Payload)
		mr.unsubscribe(msg.Topic)
	case CloseEvent:
		tr.cr.OnChannelClose(msg.Payload)
		mr.unsubscribe(msg.Topic)
	default:
		tr.cr.OnMessage(msg.Event, msg.Payload)
	}
}

func (mr *messageRouter) subscribe(topic string, cr ChannelReceiver, rr *replyRouter) {
	mr.mapLock.Lock()
	defer mr.mapLock.Unlock()
	mr.tr[topic] = &topicReceiver{cr: cr, rr: rr}
}

func (mr *messageRouter) unsubscribe(topic string) {
	mr.mapLock.Lock()
	defer mr.mapLock.Unlock()
	delete(mr.tr, topic)
}
