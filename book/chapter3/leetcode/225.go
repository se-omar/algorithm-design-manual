package main

type queue []int

func (q *queue) enqueue(x int) {
	*q = append(*q, x)
}

func (q *queue) dequeue() int {
	el := (*q)[0]
	*q = (*q)[1:]
	return el
}

func (q *queue) size() int {
	return len(*q)
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

type MyStack struct {
	q1 queue
	q2 queue
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	for this.q1.size() > 0 {
		this.q2.enqueue(this.q1.dequeue())
	}
	this.q1.enqueue(x)
	for this.q2.size() > 0 {
		this.q1.enqueue(this.q2.dequeue())
	}
}

func (this *MyStack) Pop() int {
	return this.q1.dequeue()
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	return this.q1.isEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
