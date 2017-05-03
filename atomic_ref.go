package gophoenix

import "sync/atomic"

type atomicRef struct {
	ref *int64
}

func (ic *atomicRef) nextRef() int64 {
	return atomic.AddInt64(ic.ref, 1)
}
