// Reference repo: https://github.com/afex/hystrix-go/tree/master/hystrix/rolling
// Every bucket span 1 second
package rolling_counter

import (
	"sync"
	"time"
)

// Window tracks a Bucket over a bounded number of time buckets.
// Currently the buckets are one second long and only the last 10 seconds are kept
type Window struct {
	Buckets map[int64]*Bucket
	Mutex   *sync.RWMutex
	size    int64
}

type WindowOpts struct {
	Size int64
}

type Bucket struct {
	Value float64
}

// NewWindow initializes a RollingWindow struct
func NewWindow(opts WindowOpts) *Window {
	return &Window{
		Buckets: make(map[int64]*Bucket),
		Mutex:   &sync.RWMutex{},
		size:    opts.Size,
	}
}

// Append appends val or vals as Bucket's value
// to window with key of time.Now().Unix()
func (w *Window) Append(vals ...float64) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	for _, val := range vals {
		w.Buckets[time.Now().Unix()] = &Bucket{Value: val}
		time.Sleep(1 * time.Second)
	}
	w.rmOldBuckets()
}

// CurBucket return Bucket at index of timestamp now
// or create it while it not exist.
func (w *Window) CurBucket() *Bucket {
	now := time.Now().Unix()
	if bucket, ok := w.Buckets[now]; !ok {
		bucket = &Bucket{}
		w.Buckets[now] = bucket
	}
	return w.Buckets[now]
}

// rmOldBuckets remove Bucket element in window
// while it is out off window timestamp bound
func (w *Window) rmOldBuckets() {
	// time over w.size second is out off the bound
	now := time.Now().Unix() - w.size
	for timestamp := range w.Buckets {
		if timestamp <= now {
			delete(w.Buckets, timestamp)
		}
	}
}

func (w *Window) Increment(val float64) {
	if val == 0 {
		return
	}

	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	w.CurBucket().Value += val
	w.rmOldBuckets()
}

// Max get maximum element in window since `now.Unix()-w.size`
func (w *Window) Max(now time.Time) float64 {
	var max float64

	w.Mutex.RLock()
	defer w.Mutex.RUnlock()

	for t, b := range w.Buckets {
		if t >= now.Unix()-w.size {
			if b.Value > max {
				max = b.Value
			}
		}
	}
	return max
}

// UpdateMax update the maximum value in the current busket.
func (w *Window) UpdateMax(val float64) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	b := w.CurBucket()
	if b.Value < val {
		b.Value = val
	}
	w.rmOldBuckets()
}

// Sum get sum of elements in window since `now.Unix()-w.size`
func (w *Window) Sum(now time.Time) float64 {
	sum := float64(0)

	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	for t, b := range w.Buckets {
		if t >= now.Unix()-w.size {
			sum += b.Value
		}
	}
	return sum
}

func (w *Window) Avg(now time.Time) float64 {
	return w.Sum(now) / float64(len(w.Buckets))
}
