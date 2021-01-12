package sliding_window

import (
	"container/list"
	"sync"
)

type bucket struct {
	Success   int
	Failure   int
	Timeout   int
	Rejection int
}

type Window struct {
	Size    int
	Buckets *list.List
	sync.RWMutex
}
