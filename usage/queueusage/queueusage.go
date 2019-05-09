package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/queue"
)

func main() {
	q := new(queue.Queue)
	q.Init()
	q.Enqueue(99)
	q.Enqueue(1)
	q.Enqueue(7)
	q.Enqueue(9001)
	q.Print()
	q.Dequeue()
	q.Print()
	fmt.Printf("Front of queue is %v\n", q.Peek())
	q.Dequeue()
	q.Print()
	q.Dequeue()
	q.Print()
}
