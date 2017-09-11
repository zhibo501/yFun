package ben

import (
	"fmt"
	"time"
)

const (
	LoopsPerCount = 10000
	MaxThreads    = 1000000
	CacheLine     = 128
)

type TestFun func(uint32, *uint64)

type timeStamp struct {
	begin time.Time
	end   time.Time
}

func Run(f []TestFun, pri []uint32) {
	var tpri uint32

	if len(f) != len(pri) {
		panic("function Run's para doesnt match")
	}
	for i := 0; i < len(f); i++ {
		fmt.Printf("i:%v / f:%v / pri:%v\n", i, f[i], pri[i])
		tpri += pri[i]
		if pri[i] == 0 {
			panic("pri should NOT be zero ... ")
		}
		if f[i] == nil {
			f[i] = doNothing
		}
	}
}
func checkRunPara() {

}

//count以万为单位
func RunEx(count uint32, threads uint32, f func(uint32, *uint64), desc string) {
	var i uint32

	total, loop, tnum := resize(count, threads)
	if f == nil {
		f = doNothing
	}

	tags, tmps, done := benAlloc(tnum)
	fmt.Printf("\nRun %v times in %v routines for %v \n", total, tnum, desc)

	t1 := time.Now()
	for i = 0; i < tnum; i++ {
		go func() {
			var j uint32
			var t timeStamp
			//worker begin
			t.begin = time.Now()
			for j = 0; j < loop; j++ {
				f(i, &(tags[i]))
			}
			//worker end
			t.end = time.Now()
			done <- t
		}()
	}
	for i = 0; i < tnum; i++ {
		tmps[i] = <-done
	}
	dt := time.Since(t1).Nanoseconds()

	//calc pps&latency
	begin := tmps[0].begin
	end := tmps[0].end
	dl := end.Sub(begin).Nanoseconds()
	for i = 1; i < tnum; i++ {
		dl += tmps[i].end.Sub(tmps[i].begin).Nanoseconds()

		if tmps[i].begin.Before(begin) {
			begin = tmps[i].begin
		}
		if tmps[i].end.After(end) {
			end = tmps[i].end
		}
	}
	dr := end.Sub(begin).Nanoseconds()

	pps := float64(total) * 1e9 / float64(dr)
	latency := float64(dl) / float64(total)
	fmt.Printf("time     : dt[%v] dr[%v] dl[%v] dt-dr[%v] (ns)\n", dt, dr, dl, dt-dr)
	fmt.Printf("pps      : %.0f kpps | %.2f ns/op\n", pps/1000, float64(dr)/float64(total))
	fmt.Printf("latency  : %.2f ns\n", latency)
}

func resize(count uint32, threads uint32) (uint32, uint32, uint32) {
	var k uint32

	if threads == 0 {
		threads = 1
	} else if threads > MaxThreads {
		threads = MaxThreads
	}

	k = count / threads
	if k < 1 {
		k = 1 //每线程至少1w次调用
	}
	if count%threads > 0 {
		k++
	}

	return k * LoopsPerCount * threads, k * LoopsPerCount, threads
}
func benAlloc(threadsNum uint32) ([]uint64, []timeStamp, chan timeStamp) {
	vtags := make([]uint64, threadsNum*(CacheLine/8+1))
	stamps := make([]timeStamp, threadsNum+1)
	channels := make(chan timeStamp, threadsNum+1)

	if vtags == nil || stamps == nil || channels == nil {
		panic("benchmark resoure alloc fail ... ")
	}
	return vtags, stamps, channels
}

func doNothing(uint32, *uint64) {

}
