package implementqueueusingstacks

// Accepted

type MyQueue struct {
	arr []int
}

func Constructor() MyQueue {
	return MyQueue{
		arr: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.arr = append(q.arr, x)
}

func (q *MyQueue) Pop() int {
	if len(q.arr) == 0 {
		return -1
	}

	ans := q.arr[0]

	q.arr = q.arr[1:]

	return ans
}

func (q *MyQueue) Peek() int {
	if len(q.arr) == 0 {
		return -1
	}

	return q.arr[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.arr) == 0
}
