package yFun

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type ShareData struct {
	sync.RWMutex
	data unsafe.Pointer
	obj  Dataer
}

type Dataer interface {
	Clone(unsafe.Pointer) unsafe.Pointer
	GetSelfPointer() unsafe.Pointer
}

func NewShareData(initObj Dataer) *ShareData {
	p := &ShareData{}
	p.data = initObj.GetSelfPointer()
	p.obj = initObj

	return p
}
func (o *ShareData) Update(f func(unsafe.Pointer)) {
	o.Lock()

	n := o.obj.Clone(o.data)
	f(n)
	atomic.StorePointer(&o.data, n) //本条指令需要保证原子性

	o.Unlock()
}
func (o *ShareData) Readonly(f func(unsafe.Pointer)) {
	p := atomic.LoadPointer(&o.data)
	f(p)
}
