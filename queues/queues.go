package main

import (
	"fmt"
)

//Queue ...
type Queue struct {
	items []int
}

//Enqueue ...
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue ...
func (q *Queue) Dequeue() *int {
	if len(q.items) > 0 {
		item := q.items[0]
		q.items = q.items[1:]
		return &item
	}
	return nil
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q)
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q)
}
