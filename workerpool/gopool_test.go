package pool

import (
	"sync"
	"testing"
)

var MaxData = 10000

func Gopool() {
	wg := new(sync.WaitGroup)
	data := make(chan int, 100)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			for _ = range data {

			}
		}(i)
	}

	for i := 0; i < MaxData; i++ {
		data <- i
	}

	close(data)
	wg.Wait()
}

func Nopool() {
	wg := new(sync.WaitGroup)

	for i := 0; i < MaxData; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
		}(i)
	}

	wg.Wait()
}

func BenchmarkGopool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gopool()
	}
}

func BenchmarkNopool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nopool()
	}
}
