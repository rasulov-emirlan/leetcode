package semaphore

type Semaphore struct {
	permits chan struct{}
}

func New(capacity int) Semaphore {
	if capacity <= 0 {
		panic("semaphore capacity must be > 0")
	}
	return Semaphore{permits: make(chan struct{}, capacity)}
}

func (s *Semaphore) Acquire() {
	s.permits <- struct{}{}
}

func (s *Semaphore) Release() {
	select {
	case <-s.permits:
		return
	default:
		panic("semaphore: release called without a matching acquire")
	}
}
