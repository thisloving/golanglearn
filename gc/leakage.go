package main

import (
	"runtime"
	"time"
)

var MaxLoop = 100000000

type Node struct {
	next    *Node
	payload [4]byte
}

//escape
func f2() *Node {
	curr := new(Node)
	for i := 0; i < MaxLoop; i++ {
		curr.next = new(Node)
		curr = curr.next
	}

	return curr
}

func f1() {
	curr := new(Node)
	for i := 0; i < MaxLoop; i++ {
		curr.next = new(Node)
		curr = curr.next
	}
}

func main() {
	//f1()
	f2() //escape
	time.Sleep(time.Second * 30)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	println(memStats.HeapInuse)
}
