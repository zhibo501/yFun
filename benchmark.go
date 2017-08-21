package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	ben(100, 1, atomic_plus)
	atomic_print()
	ben(100, 2, atomic_plus)
	atomic_print()
}

//count以万为单位
func ben(count uint32, threads uint32, f func()) {
	var i, k uint32

	k = count / threads
	if k < 1 {
		k = 1 //每线程至少1w次调用
	}

	done := make(chan bool)

	//begin
	t1 := time.Now()
	for i = 0; i < threads; i++ {
		go func() {
			var j uint32
			for j = 0; j < k*10000; j++ {
				f()
			}
			done <- true
		}()
	}

	for i = 0; i < threads; i++ {
		<-done
	}
	//end
	d := time.Since(t1).Nanoseconds()
	total := k * 10000 * threads
	op := float64(d) / float64(total)
	pps := float64(total) * 1e9 / float64(d)
	fmt.Println("t1 -> t2 :", d, "ns")
	fmt.Println("total    :", total)
	fmt.Println("ns/op    :", op, "ns/op")
	fmt.Println("pps      :", pps/1000, "kpps")
}

var g_atomic uint64

func atomic_plus() {
	g_atomic++
}
func atomic_print() {
	fmt.Println("g_atomic : ", g_atomic)
}
