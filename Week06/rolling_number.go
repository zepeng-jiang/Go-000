package Week06

import (
	"sync"
	"sync/atomic"
	"time"
)

// EventTypeAndNumber 事件类型及其对应的数量
type EventTypeAndNumber struct {
	// 成功数
	success int32
	// 失败数
	fail int32
}

// Bucket 桶
type Bucket struct {
	data        EventTypeAndNumber
	windowStart int64
}

// RollingNumber .
type RollingNumber struct {
	buckets []*Bucket
	size    int64
	width   int64
	tail    int64
	mux     sync.RWMutex
}

func NewRollingNumber(size, width int64) *RollingNumber {
	return &RollingNumber{
		size:    size,
		width:   width,
		buckets: make([]*Bucket, size),
		tail:    0,
	}
}

func (rn *RollingNumber) GetCurrent() *Bucket {
	rn.mux.Lock()
	defer rn.mux.Unlock()

	current := time.Now().Unix()
	if rn.tail == 0 && rn.buckets[rn.tail] == nil {
		bk := &Bucket{
			data:        EventTypeAndNumber{},
			windowStart: current,
		}
		rn.buckets[rn.tail] = bk
		return bk
	}

	last := rn.buckets[rn.tail]
	if current < last.windowStart+rn.width {
		return last
	}

	for i := 0; i < int(rn.size); i++ {
		last := rn.buckets[rn.tail]
		if current < last.windowStart+rn.width {
			return last
		} else if current-(last.windowStart+rn.width) > rn.size*rn.width {
			rn.tail = 0
			rn.buckets = make([]*Bucket, rn.size)
			rn.mux.Unlock()
			return rn.GetCurrent()
		} else {
			rn.tail++
			bk := &Bucket{
				data:        EventTypeAndNumber{},
				windowStart: last.windowStart + rn.width,
			}

			if rn.tail >= rn.size {
				copy(rn.buckets[:], rn.buckets[1:])
				rn.tail--
			}
			rn.buckets[rn.tail] = bk
		}
	}
	return rn.buckets[rn.tail]
}

func (rn *RollingNumber) IncrSuccess() {
	bk := rn.GetCurrent()
	atomic.AddInt32(&bk.data.success, 1)
}

func (rn *RollingNumber) IncrFail() {
	bk := rn.GetCurrent()
	atomic.AddInt32(&bk.data.fail, 1)
}

func (rn *RollingNumber) GetRollingSum() *EventTypeAndNumber {
	en := &EventTypeAndNumber{}
	rn.mux.RLock()
	defer rn.mux.RUnlock()
	for _, v := range rn.buckets {
		if nil == v {
			continue
		}
		en.success += v.data.success
		en.fail += v.data.fail
	}
	return en
}
