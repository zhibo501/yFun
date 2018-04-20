package queue

import (
	"testing"
)

func TestUniqListFunc(t *testing.T) {
	l := NewUniqList()

	// push one, pop one
	l.Push("no1", "no1")
	if 1 != l.Count() {
		t.Errorf("error count .... %v != 1 ", l.Count())
	}
	v := l.Pop()
	if v.Key != "no1" {
		t.Errorf("error .... %v ", v)
	}
	if 0 != l.Count() {
		t.Errorf("error count .... %v != 0 ", l.Count())
	}

	// push n, pop n-1
	l.Push("no1", "no1")
	l.Push("no2", "no2")
	l.Push("no3", "no3")
	l.Push("no1", "no1")
	l.Push("no2", "no2")
	l.Push("no3", "no3")
	l.Show()

	if 3 != l.Count() {
		t.Errorf("error count .... %v != 3 ", l.Count())
	}
	v = l.Pop()
	if v.Key != "no1" {
		t.Errorf("error .... %v ", v)
	}
	if 2 != l.Count() {
		t.Errorf("error count .... %v != 2 ", l.Count())
	}
	l.Show()
	v = l.Pop()
	if v.Key != "no2" {
		t.Errorf("error .... %v ", v)
	}
	if 1 != l.Count() {
		t.Errorf("error count .... %v != 1 ", l.Count())
	}
	l.Show()
}

func Benchmark_TestUniqList(b *testing.B) {
	l := NewUniqList()
	l.Push("no1", "no1")
	l.Push("no2", "no2")
	l.Push("no3", "no3")

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Push("non", "non")
			_ = l.Pop()
		}
	})
}
