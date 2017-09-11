package main

import (
	"fmt"
	ben "github.com/zhibo501/yFun/ben"
	"sync"
)

func main() {
	/*
		ben.RunEx(100, 1, simple_plus, "simple_plus")
		ben.RunEx(100, 2, simple_plus, "simple_plus")
		ben.RunEx(100, 4, simple_plus, "simple_plus")
		ben.RunEx(100, 10, simple_plus, "simple_plus")
		ben.RunEx(100, 20, simple_plus, "simple_plus")
		ben.RunEx(100, 40, simple_plus, "simple_plus")
		ben.RunEx(100, 100, simple_plus, "simple_plus")

		ben.RunEx(100, 1, atomic_plus, "atomic_plus")
		ben.RunEx(100, 2, atomic_plus, "atomic_plus")
		ben.RunEx(100, 4, atomic_plus, "atomic_plus")
		ben.RunEx(100, 10, atomic_plus, "atomic_plus")
		ben.RunEx(100, 20, atomic_plus, "atomic_plus")
		ben.RunEx(100, 40, atomic_plus, "atomic_plus")
		ben.RunEx(100, 100, atomic_plus, "atomic_plus")
	*/
	ben.Run([]ben.TestFun{simple_plus, atomic_plus}, []uint32{50, 50})

	ben.RunEx(100, 1, lock_set, "lock_set")
	ben.RunEx(100, 2, lock_set, "lock_set")
	ben.RunEx(100, 4, lock_set, "lock_set")
}

////////////////////////////////////////////////
// simple plus
////////////////////////////////////////////////
func simple_plus(id uint32, v *uint64) {
	var i uint32
	i++
	//*v++
}

////////////////////////////////////////////////
// atomic plus
////////////////////////////////////////////////
var g_atomic uint64

func atomic_plus(id uint32, v *uint64) {
	g_atomic++
}
func atomic_print() {
	fmt.Println("g_atomic : ", g_atomic)
}

//
// public struct
//
type A struct {
	name string
	d1   uint32
	d2   uint32
	buf  [256]byte
}

var gVar A

////////////////////////////////////////////////
// lock set
////////////////////////////////////////////////
var mu sync.Mutex

func lock_set(id uint32, v *uint64) {
	mu.Lock()
	gVar.name = "Concurrency"
	gVar.d1 = id
	gVar.d2 = id
	gVar.buf[0] = 'a'
	gVar.buf[1] = 'b'
	gVar.buf[2] = 'c'
	gVar.buf[3] = 'd'
	copy(gVar.buf[4:], "he012345678901234567890123456789")
	copy(gVar.buf[32:], "he012345678901234567890123456789")
	mu.Unlock()
}

type DB struct {
	db *Connect
}

func (d *DB) Write(fn func(*DB)) *DB {
	v := d.clone()
	fn(v)
	return v
}
func (d *DB) Read(fn func(*DB)) *DB {
	fn(d)
}
