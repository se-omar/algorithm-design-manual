package main

type stack []int

func (s *stack) pop() int {
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) push(el int) {
	*s = append(*s, el)
}

func (s *stack) size() int {
	return len(*s)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

type MyQueue struct {
	s1 stack
	s2 stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	if this.s1.isEmpty() {
		this.s1.push(x)
	} else {
		for this.s1.size() > 0 {
			this.s2.push(this.s1.pop())
		}
		this.s1.push(x)
		for this.s2.size() > 0 {
			this.s1.push(this.s2.pop())
		}
	}
}

func (this *MyQueue) Pop() int {
	return this.s1.pop()
}

func (this *MyQueue) Peek() int {
	return this.s1.peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.isEmpty()
}

/**
* Your MyQueue object will be instantiated and called as such:
obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Peek();
* param_4 := obj.Empty();
*/
