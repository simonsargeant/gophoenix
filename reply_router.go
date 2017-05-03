package gophoenix

import "sync"

type replyRouter struct {
	mapLock sync.RWMutex
	rr map[int64]replyCallback
}

type replyCallback func (payload interface{})

func newReplyRouter() *replyRouter {
	return &replyRouter{
		rr: make(map[int64]replyCallback),
	}
}

func (rr *replyRouter) routeReply(msg *Message) {
	rr.mapLock.RLock()
	rc, ok := rr.rr[msg.Ref]
	rr.mapLock.RUnlock()

	if !ok {
		return
	}

	rc(msg.Payload)
}

func (rr *replyRouter) subscribe(ref int64, callback replyCallback) {
	rr.mapLock.Lock()
	defer rr.mapLock.Unlock()
	rr.rr[ref] = callback
}
