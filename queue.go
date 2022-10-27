package main

import "fmt"

func enqueue(q []int, val int) []int {
	q = append(q, val)
	fmt.Println("enqued vslue is : ", val)

	return q
}

func dequeue(q []int) []int {
	deq := q[0]
	fmt.Println("Dequed value: ", deq)
	return q[1:]
}

func main() {
	var queue []int // Make a queue of ints.

	queue = enqueue(queue, 10)
	queue = enqueue(queue, 20)
	queue = enqueue(queue, 30)

	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	queue = enqueue(queue, 40)
	fmt.Println("Queue:", queue)
}
