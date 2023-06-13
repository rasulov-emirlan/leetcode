package minstack

type MinStack struct {
	top *StackNode
	min int
}

type StackNode struct {
	data    int
	next    *StackNode
	lastmin int
}

func Constructor() MinStack {
	return MinStack{}
}

func (minStack *MinStack) Push(val int) {
	newtop := &StackNode{data: val, next: minStack.top}

	if minStack.top == nil {
		minStack.min = val
	} else {
		newtop.lastmin = minStack.min
	}

	minStack.top = newtop
	if minStack.top.data < minStack.min {
		minStack.min = minStack.top.data
	}
}

func (minStack *MinStack) Pop() {
	if minStack.top.next == nil {
		minStack.top = nil
		return
	}
	minStack.min = minStack.top.lastmin
	minStack.top = minStack.top.next
}

func (minStack *MinStack) Top() int {
	return minStack.top.data
}

func (minStack *MinStack) GetMin() int {
	return minStack.min
}
