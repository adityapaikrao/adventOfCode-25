package queue

import (
	"container/list"
	"fmt"
)

type Queue struct {
	data *list.List
}

/*Returns a New Qeueu*/
func NewQueue() *Queue {
	data := list.New()
	return &Queue{data: data}
}

/*Push an Element to the queue*/
func (q *Queue) Push(value any) {
	q.data.PushBack(value)
}

/*Pop the front of the Queue*/
func (q *Queue) Popleft() any {
	front := q.data.Front()
	q.data.Remove(front)

	return front.Value
}

/*Get Element at Front of the Queue*/
func (q *Queue) Front() any {
	front := q.data.Front()

	return front.Value
}

/*Check if Queue is Empty*/
func (q *Queue) IsEmpty() bool {
	return q.data.Len() == 0
}

/*Get size of the queue*/
func (q *Queue) Size() int {
	return q.data.Len()
}
func (q *Queue) PrintQueue() {
	for e := q.data.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
