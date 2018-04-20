package queue

import (
	"fmt"
	"sync"
	//"sync/atomic"
)

type CList struct {
	head *lNode
	tail *lNode
	num  int32
	sync.RWMutex
}

func NewCList() *CList {
	n := &lNode{}
	return &CList{head: n, tail: n, num: 0}
}

func (l *CList) Push(k string, v interface{}) {
	n := &lNode{data: &Value{k, v}}

	l.Lock()
	cur := l.tail
	cur.next = n
	n.prev = cur
	l.tail = n
	//atomic.AddInt32(&l.num, 1)
	l.num++
	l.Unlock()
}

func (l *CList) Pop() *Value {
	l.Lock()
	if l.head == l.tail {
		l.Unlock()
		return nil
	}
	cur := l.head.next
	cur.prev = nil
	l.head = cur
	//atomic.AddInt32(&l.num, -1)
	l.num--
	l.Unlock()

	return cur.data
}

func (l *CList) Count() int32 {
	l.RLock()
	n := l.num
	l.RUnlock()
	return n
}

func (l *CList) Show() {
	l.RLock()
	fmt.Printf("Total %v nodes ...\n", l.num)
	for n := l.head.next; n != nil; n = n.next {
		fmt.Printf("  lNode : {%v} \n", n.data)
	}
	fmt.Printf("  ---------------\n")
	for n := l.tail; n != l.head; n = n.prev {
		fmt.Printf("  lNode : {%v} \n", n.data)
	}
	l.RUnlock()
}
