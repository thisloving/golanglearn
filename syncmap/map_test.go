package syncmap

import (
	"sync"
	"testing"
)

func TestStore(t *testing.T) {
	m := &Map{}
	// write to dirty
	m.Store(1, nil)
	t.Logf("after store %d", m.Len())

	// prompt dirty to read
	m.Load(2)
	m.Load(3)

	m.Delete(1)
	t.Logf("after delete %d", m.Len())

	m.Store(1, nil)
	t.Logf("after store %d", m.Len())

	if m.Len() != 1 {
		t.Fatalf("%d", m.Len())
	}
}

func TestLen(t *testing.T) {
	m := &Map{}
	var wg sync.WaitGroup

	num := 100000

	for i := 0; i < num; i++ {
		wg.Add(1)

		go func(a int) {
			m.Store(a, nil)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2*num; i++ {
		wg.Add(1)

		go func(a int) {
			m.Load(a)
			//m.LoadOrStore(a, nil)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// t.Logf("%d", m.Len())
	t.Logf("%d", m.Size())

	for i := 0; i < 2*num; i++ {
		wg.Add(1)

		go func(a int) {
			m.Delete(a)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// m.Range(func(k, v interface{}) bool {
	// 	m.Delete(k)
	// 	return true
	// })

	t.Logf("%d", m.Size())
	t.Logf("%d", m.Len())

	if m.Len() != m.Size() {
		t.Fatalf("%d", m.Len())
	}
}
