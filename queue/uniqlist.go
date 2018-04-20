package queue

import (
	"fmt"
	"sync"
	//"sync/atomic"
)

type UniqList struct {
	head *lNode
	tail *lNode
	num  int32
	m    map[string]interface{}
	sync.RWMutex
}

func NewUniqList() *UniqList {
	n := &lNode{}
	return &UniqList{head: n, tail: n, num: 0, m: make(map[string]interface{}, 128)}
}

func (l *UniqList) Push(k string, v interface{}) {
	n := &lNode{data: &Value{k, v}}

	l.Lock()
	_, ok := l.m[k]
	if ok {
		l.Unlock()
		return
	}
	cur := l.tail
	cur.next = n
	n.prev = cur
	l.tail = n
	l.m[k] = n
	//atomic.AddInt32(&l.num, 1)
	l.num++
	l.Unlock()
}

func (l *UniqList) Pop() *Value {
	l.Lock()
	if l.head == l.tail {
		l.Unlock()
		return nil
	}
	cur := l.head.next
	cur.prev = nil
	l.head = cur
	delete(l.m, cur.data.Key)
	//atomic.AddInt32(&l.num, -1)
	l.num--
	l.Unlock()

	return cur.data
}

func (l *UniqList) Count() int32 {
	l.RLock()
	n := l.num
	l.RUnlock()
	return n
}

func (l *UniqList) Show() {
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
