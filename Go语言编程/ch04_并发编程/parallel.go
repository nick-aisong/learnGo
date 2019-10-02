// parallel
package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

type Vector []float64

// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了
}

const NCPU = 16 // 假设总共有16核
func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// 等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	// 到这里表示所有计算已经结束
}

func (v Vector) Op(i float64) float64 {
	time.Sleep(time.Second)
	return math.Sqrt(i * i * i)
}

func main() {

	runtime.GOMAXPROCS(16)

	var a, v Vector

	for i := 0; i <= 10000000; i++ {
		v = append(v, rand.Float64())
	}

	fmt.Println(runtime.NumCPU())
	//fmt.Println(v)

	b := time.Now()
	a.DoAll(v)
	// 没用，直接开了协程序，程序直接跑完
	e := time.Now()
	fmt.Print(b)
	fmt.Println()
	fmt.Print(e)
	fmt.Println()

}

// 如果i过大，10000000000
// runtime: VirtualAlloc of 2410020864 bytes failed with errno=1455
// fatal error: out of memory

// runtime stack:
// runtime.throw(0x4c9893, 0xd)
//         C:/Go/src/runtime/panic.go:616 +0x88
// runtime.sysMap(0xc174bc0000, 0x8fa60000, 0xc042024001, 0x5625f8)
//         C:/Go/src/runtime/mem_windows.go:122 +0x13b
// runtime.(*mheap).sysAlloc(0x549c20, 0x8fa60000, 0x44b700)
//         C:/Go/src/runtime/malloc.go:470 +0xdb
// runtime.(*mheap).grow(0x549c20, 0x47d2b, 0x0)
//         C:/Go/src/runtime/mheap.go:907 +0x67
// runtime.(*mheap).allocSpanLocked(0x549c20, 0x47d2b, 0x562608, 0x25efe38)
//         C:/Go/src/runtime/mheap.go:820 +0x308
// runtime.(*mheap).alloc_m(0x549c20, 0x47d2b, 0xffffffffffff0101, 0x25efe70)
//         C:/Go/src/runtime/mheap.go:686 +0x126
// runtime.(*mheap).alloc.func1()
//         C:/Go/src/runtime/mheap.go:753 +0x54
// runtime.(*mheap).alloc(0x549c20, 0x47d2b, 0x2000101, 0x8f4b45fe5)
//         C:/Go/src/runtime/mheap.go:752 +0x91
// runtime.largeAlloc(0x8fa56000, 0xffffffffffff0100, 0x8f4b45fe5)
//         C:/Go/src/runtime/malloc.go:826 +0x9b
// runtime.mallocgc.func1()
//         C:/Go/src/runtime/malloc.go:721 +0x4d
// runtime.systemstack(0x0)
//         C:/Go/src/runtime/asm_amd64.s:409 +0x7e
// runtime.mstart()
//         C:/Go/src/runtime/proc.go:1175

// goroutine 1 [running]:
// runtime.systemstack_switch()
//         C:/Go/src/runtime/asm_amd64.s:363 fp=0xc04242fd58 sp=0xc04242fd50 pc=0x4
// 4c810
// runtime.mallocgc(0x8fa56000, 0x0, 0x15c8bc8b33e93300, 0xc04242fe30)
//         C:/Go/src/runtime/malloc.go:720 +0x8e1 fp=0xc04242fdf8 sp=0xc04242fd58 p
// c=0x40e8d1
// runtime.growslice(0x4a5b60, 0xc09e322000, 0xe5d5400, 0xe5d5400, 0xe5d5401, 0xc09
// e322000, 0xb7ddc00, 0xe5d5400)
//         C:/Go/src/runtime/slice.go:172 +0x224 fp=0xc04242fe60 sp=0xc04242fdf8 pc
// =0x439624
// main.main()
//         C:/Users/MY/Documents/GitHub/learnGo/Go语言编程/ch04_并发编程/parallel.g
// o:43 +0xdf fp=0xc04242ff88 sp=0xc04242fe60 pc=0x4955df
// runtime.main()
//         C:/Go/src/runtime/proc.go:198 +0x20e fp=0xc04242ffe0 sp=0xc04242ff88 pc=
// 0x42a29e
// runtime.goexit()
//         C:/Go/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc04242ffe8 sp=0xc04242ffe0
// pc=0x44ed31
// exit status 2
