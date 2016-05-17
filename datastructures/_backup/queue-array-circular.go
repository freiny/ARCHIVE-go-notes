package main

import (
	"fmt"
	"math"
)

func main() {

	q := queue{}.new()
	// OUTPUT: {3 0 0 0 [0 0 0]}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	// OUTPUT: {3 1 0 1 [1 0 0]}
	// OUTPUT: {3 2 0 2 [1 2 0]}
	// OUTPUT: {3 3 0 0 [1 2 3]}

	// q.enqueue(4)
	// // OUTPUT: ERROR: queue size exceeded

	q.dequeue()
	q.enqueue(4)
	// OUTPUT: 1
	// OUTPUT: {3 2 1 0 [1 2 3]}
	// OUTPUT: {3 3 1 1 [4 2 3]}

	q.dequeue()
	q.enqueue(5)
	// OUTPUT: 2
	// OUTPUT: {3 2 2 1 [4 2 3]}
	// OUTPUT: {3 3 2 2 [4 5 3]}

	q.dequeue()
	q.enqueue(6)
	// OUTPUT: 3
	// OUTPUT: {3 2 0 2 [4 5 3]}
	// OUTPUT: {3 3 0 0 [4 5 6]}

	q.dequeue()
	q.dequeue()
	q.dequeue()
	// OUTPUT: 4
	// OUTPUT: {3 2 1 0 [4 5 6]}
	// OUTPUT: 5
	// OUTPUT: {3 1 2 0 [4 5 6]}
	// OUTPUT: 6
	// OUTPUT: {3 0 0 0 [4 5 6]}

	// q.dequeue()
	// // OUTPUT: ERROR: queue size exceeded

}

const max = 3

type queue struct {
	sizeMax int
	size    int
	head    int
	tail    int
	list    [max]int
}

func (q queue) new() queue {
	q.sizeMax = len(q.list)
	q.print()
	return q
}

func (q *queue) enqueue(item int) {
	if q.size < q.sizeMax {
		q.list[q.tail] = item
		q.tail = int(math.Mod(float64(q.tail+1), float64(q.sizeMax)))
		q.size++
		q.print()
	} else {
		fmt.Println("ERROR: queue size exceeded")
	}
}

func (q *queue) dequeue() int {
	if q.size > 0 {
		ret := q.list[q.head]
		q.head = int(math.Mod(float64(q.head+1), float64(q.sizeMax)))
		q.size--
		fmt.Println(ret)
		q.print()
		return ret
	} else {
		fmt.Println("ERROR: dequeue failed no more items in list")
		return 0
	}
}

func (q queue) print() {
	fmt.Println(q)
}
