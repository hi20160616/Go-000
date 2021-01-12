package demo

import (
	"math/rand"
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sw := NewSlidingWindow(100)
	var r int
	for i := 0; i < 1000; i++ {
		r = rand.Intn(3)
		if r == 1 {
			sw.AddSuccess()
		} else {
			sw.AddFail()
		}
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	}
	t.Log("1 sec bucket len: ", sw.Len())
	for _, item := range sw.Data(3) {
		t.Log(item.success, item.fail)
	}
	t.Log("===================")
	for _, item := range sw.Data(5) {
		t.Log(item.success, item.fail)
	}
}
