package gophoenix

import "sync/atomic"

type atomic_ref struct {
	ref *int64
}

func (ic *atomic_ref) nextRef() int64 {
	return atomic.AddInt64(ic.ref, 1)
}
