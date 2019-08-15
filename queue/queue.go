package queue

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/linkedlist"
)

//need to work on trying to remove if queue is empty
type Queue struct {
	List *linkedlist.LinkedList
	End  *linkedlist.Node
}

//init queue
func (q *Queue) Init() {
	q.List = new(linkedlist.LinkedList)
	q.List.Head = new(linkedlist.Node)
	q.End = nil
}

//set the end pointer
func (q *Queue) SetEnd() {
	p := new(linkedlist.Node)
	p = q.List.Head
	if p.Data == nil && p.Next == nil {
		//empty queue
		fmt.Println("empty queue, setting end")
		q.End = p
	} else {
		for p.Next != nil {
			p = p.Next
		}
		q.End = p
	}

}

//add to end of queue. uses pointer to end for time efficiency
func (q *Queue) Enqueue(v interface{}) {
	if q.End == nil {
		fmt.Println("Enqueuing, setting end first")
		q.SetEnd()
		q.End.Data = v
	} else {
		p := new(linkedlist.Node)
		p.Data = v
		q.End.Next = p
		q.End = p
	}
}

//remove value from front of queue
func (q *Queue) Dequeue() interface{} {
	value := q.List.Head.Data
	q.List.RemoveStart()
	return value
}

//peek at front of queue
func (q *Queue) Peek() interface{} {
	return q.List.Head.Data
}

//traverse list
func (q *Queue) Print() {
	q.List.Traverse()
}
