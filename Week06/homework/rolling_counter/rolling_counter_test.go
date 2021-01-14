package rolling_counter

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMax(t *testing.T) {
	Convey("When adding values to a rolling window", t, func() {
		opts := WindowOpts{Size: 10}
		w := NewWindow(opts)
		for _, x := range []float64{10, 11, 9} {
			w.Append(x)
		}

		Convey("It should know the maximum", func() {
			So(w.Max(time.Now()), ShouldEqual, 11)
		})
	})
}

func TestAvg(t *testing.T) {
	Convey("When adding values to a rolling window", t, func() {
		opts := WindowOpts{Size: 10}
		w := NewWindow(opts)
		for _, x := range []float64{0.5, 1.5, 2.5, 3.5, 4.5} {
			w.Append(x)
		}
		Convey("It should know the average", func() {
			So(w.Avg(time.Now()), ShouldEqual, 2.5)
		})
	})
}

func BenchmarkNewWindowIncrement(b *testing.B) {
	b.ResetTimer()
	opts := WindowOpts{Size: 10}
	for i := 0; i < b.N; i++ {
		NewWindow(opts).Increment(1)
	}
}

func BenchmarkNewWindowUpdateMax(b *testing.B) {
	b.ResetTimer()
	opts := WindowOpts{Size: 10}
	for i := 0; i < b.N; i++ {
		NewWindow(opts).UpdateMax(float64(i))
	}
}
