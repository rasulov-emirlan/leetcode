package semaphore

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSemaphoreAcquireBlocksUntilRelease(t *testing.T) {
	sem := New(1)
	sem.Acquire() // fill capacity

	unblocked := make(chan struct{})
	go func() {
		sem.Acquire() // should block until release
		close(unblocked)
	}()

	select {
	case <-unblocked:
		t.Fatalf("Acquire should be blocked before release")
	case <-time.After(50 * time.Millisecond):
		// expected: still blocked
	}

	sem.Release() // free one permit

	select {
	case <-unblocked:
		// success: goroutine proceeded after release
	case <-time.After(1 * time.Second):
		t.Fatalf("Acquire did not unblock after release")
	}
}

func TestReleaseWithoutAcquirePanics(t *testing.T) {
	sem := New(1)

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		sem.Release()
	}()
	if !didPanic {
		t.Fatalf("expected panic when releasing without prior acquire")
	}
}

func TestNewWithInvalidCapacityPanics(t *testing.T) {
	for _, cap := range []int{0, -1, -10} {
		didPanic := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()
			_ = New(cap)
		}()
		if !didPanic {
			t.Fatalf("expected panic for capacity %d", cap)
		}
	}
}

func TestConcurrentAcquireRelease(t *testing.T) {
	sem := New(3)
	var wg sync.WaitGroup
	const workers = 10
	const iterations = 100

	active := 0
	var mu sync.Mutex
	maxActive := 0

	worker := func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			sem.Acquire()
			mu.Lock()
			active++
			if active > maxActive {
				maxActive = active
			}
			mu.Unlock()

			// Simulate small work
			runtime.Gosched()

			mu.Lock()
			active--
			mu.Unlock()
			sem.Release()
		}
	}

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker()
	}
	wg.Wait()

	if maxActive > 3 {
		t.Fatalf("observed concurrency %d exceeds capacity", maxActive)
	}
}
