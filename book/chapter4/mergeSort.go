package chapter4

type queue []int

func (q *queue) enqueue(x int) {
	*q = append(*q, x)
}

func (q *queue) head() int {
	return (*q)[0]
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

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	mid := len(s) / 2
	leftArr := MergeSort(s[:mid])
	rightArr := MergeSort(s[mid:])

	return merge(leftArr, rightArr)
}

func merge(left, right []int) []int {
	ql, qr := queue{}, queue{}
	total := []int{}

	for _, v := range left {
		ql.enqueue(v)
	}

	for _, v := range right {
		qr.enqueue(v)
	}

	for ql.size() > 0 && qr.size() > 0 {
		if ql.head() < qr.head() {
			total = append(total, ql.dequeue())
		} else {
			total = append(total, qr.dequeue())
		}
	}

	for ql.size() > 0 {
		total = append(total, ql.dequeue())
	}

	for qr.size() > 0 {
		total = append(total, qr.dequeue())
	}

	return total
}
