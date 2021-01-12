// Reference: https://github.com/jikebang/Go-000/blob/main/Week06/work/sliding_window.go
package demo

import (
	"container/list"
	"log"
	"sync"
	"time"
)

const (
	typeSuccess int = 1
	typeFail    int = 2
)

type metrics struct {
	success int64
	fail    int64
}

type SlidingWindow struct {
	bucket int                // bucket count
	curKey int64              // current key
	m      map[int64]*metrics // temp variable
	data   *list.List         // bucket
	sync.RWMutex
}

func NewSlidingWindow(bucket int) *SlidingWindow {
	sw := &SlidingWindow{}
	sw.bucket = bucket
	sw.data = list.New()
	return sw
}

func (sw *SlidingWindow) AddSuccess() {
	sw.incr(typeSuccess)
}

func (sw *SlidingWindow) AddFail() {
	sw.incr(typeFail)
}

func (sw *SlidingWindow) incr(t int) {
	sw.Lock()
	defer sw.Unlock()
	// init sw.m[now]
	now := time.Now().Unix()
	if _, ok := sw.m[now]; !ok {
		sw.m = make(map[int64]*metrics)
		sw.m[now] = &metrics{}
	}
	if sw.curKey == 0 { // curKey == 0 at first time.
		sw.curKey = now
	}
	if sw.curKey != now {
		sw.data.PushBack(sw.m[now]) // push new sw.m to sw.data
		// why delete sw.m? just for temp use?
		delete(sw.m, sw.curKey) // sw.m[sw.curKey] is out of date, del it.
		sw.curKey = now         // update curKey
		if sw.data.Len() > sw.bucket {
			// rm front element while data over bucket capacity
			for i := 0; i < sw.data.Len()-sw.bucket; i++ {
				sw.data.Remove(sw.data.Front())
			}
		}
	}

	switch t {
	case typeSuccess:
		sw.m[now].success++
	case typeFail:
		sw.m[now].fail++
	default:
		log.Fatal("error type")
	}
}

func (sw *SlidingWindow) Len() int {
	return sw.data.Len()
}

func (sw *SlidingWindow) Data(space int) []*metrics {
	sw.RLock()
	defer sw.RUnlock()
	var data []*metrics
	var num = 0
	var m = &metrics{}
	for i := sw.data.Front(); i != nil; i = i.Next() {
		one := i.Value.(*metrics)
		m.success += one.success
		m.fail += one.fail
		if num%space == 0 {
			data = append(data, m)
			m = &metrics{} // reset m
		}
		num++
	}
	return data
}
