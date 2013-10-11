package gobuf

import (
    "runtime"
    "sync/atomic"
)

const (
    size uint64 = 8192
    mask uint64 = size - 1
)

type Buf struct {
    items [size]interface{}
    index uint64
    nextIndex uint64
    getIndex uint64
}

func NewBuf() *Buf {
    return &Buf{index: 0, nextIndex: 1, getIndex: 1}
}

func (buf *Buf) Add(item interface{}) {
    idx := atomic.AddUint64(&buf.nextIndex, 1) - 1
    for idx > (buf.getIndex + size - 2) {
        // Allow other things to run until caught up
        runtime.Gosched()
    }
    
    // Add item to bufer
    buf.items[idx & mask] = item

    // Update index to next position
    for !atomic.CompareAndSwapUint64(&buf.index, idx - 1, idx) {
        runtime.Gosched()
    }
}

func (buf *Buf) Get() interface{} {
    idx := atomic.AddUint64(&buf.getIndex, 1) - 1
    for idx > buf.index {
        // Allow other things to run until caught up
        runtime.Gosched()
    }
    return buf.items[idx & mask]
}
