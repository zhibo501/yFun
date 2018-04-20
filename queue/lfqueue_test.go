package queue

import (
	"testing"
)

func Test_Lfque(t *testing.T) {
	q := NewLockfreeQueue()

	// push one, pop one
	q.Push("no1")
	v := q.Pop()
	if v != "no1" {
		t.Errorf("error .... %v ", v)
	}
	v = q.Pop()
	if v != nil {
		t.Errorf("error .... %v ", v)
	}

	// push n, pop n-1
	q.Push("no1")
	q.Push("no2")
	q.Push("no3")
	v = q.Pop()
	if v != "no1" {
		t.Errorf("error .... %v ", v)
	}
	v = q.Pop()
	if v != "no2" {
		t.Errorf("error .... %v ", v)
	}
	v = q.Pop()
	if v != "no3" {
		t.Errorf("error .... %v ", v)
	}
}

func Benchmark_LfqueTest(b *testing.B) {
	q := NewLockfreeQueue()
	q.Push("no1")
	q.Push("no2")

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Push("non")
			_ = q.Pop()
		}
	})
}
