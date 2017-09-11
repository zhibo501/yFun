package yFun

import (
	"testing"
	"unsafe"
)

type A struct {
	s   string
	d1  int
	d2  int
	buf [64]byte
}

func (a *A) Clone(p unsafe.Pointer) unsafe.Pointer {
	tmp := *a
	return unsafe.Pointer(&tmp)
}
func (a *A) GetSelfPointer() unsafe.Pointer {
	return unsafe.Pointer(a)
}

func AReader(p unsafe.Pointer) {
	a := (*A)(p)

	b := a.buf[0]
	b += a.buf[1]
	b += a.buf[2]
	b += a.buf[3]
}
func AWriter(p unsafe.Pointer) {
	a := (*A)(p)

	a.buf[0] = 'y'
	a.buf[1] = 'f'
	a.buf[2] = 'u'
	a.buf[3] = 'n'
	copy(a.buf[4:], "hello01234567890123456789")
	copy(a.buf[32:], "hello01234567890123456789")
}

func Benchmark_ShareDataParallel_Read(b *testing.B) {
	a := &A{}
	s := NewShareData(a)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Readonly(AReader)
		}
	})
}
func Benchmark_ShareDataParallel_Write(b *testing.B) {
	a := &A{}
	s := NewShareData(a)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Update(AWriter)
		}
	})
}
func Benchmark_ShareDataParallel_ReadAndWrite(b *testing.B) {
	a := &A{}
	s := NewShareData(a)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Readonly(AReader)
			s.Update(AWriter)
		}
	})
}
