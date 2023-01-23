package implementqueueusingstacks

import "testing"

func TestQueue(t *testing.T) {
	q := Constructor()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Peek() != 1 {
		t.Errorf("q.Peek() = %v, want %v", q.Peek(), 1)
	}

	if q.Pop() != 1 {
		t.Errorf("q.Pop() = %v, want %v", q.Pop(), 1)
	}

	if q.Peek() != 2 {
		t.Errorf("q.Peek() = %v, want %v", q.Peek(), 2)
	}

	if q.Pop() != 2 {
		t.Errorf("q.Pop() = %v, want %v", q.Pop(), 2)
	}

	if q.Peek() != 3 {
		t.Errorf("q.Peek() = %v, want %v", q.Peek(), 3)
	}

	if q.Pop() != 3 {
		t.Errorf("q.Pop() = %v, want %v", q.Pop(), 3)
	}

	if q.Empty() != true {
		t.Errorf("q.Empty() = %v, want %v", q.Empty(), true)
	}
}
