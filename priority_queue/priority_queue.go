package main

import (
	"errors"
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

//IsEmpty ...
func (q Queue) IsEmpty() bool {
	return len(q.items) == 0
}

//QueueManager ...
type QueueManager struct {
	PriorityQueues []Queue
}

//NewQueueManager ...
func NewQueueManager() QueueManager {
	return QueueManager{
		PriorityQueues: []Queue{
			Queue{},
			Queue{},
			Queue{},
			Queue{},
			Queue{},
		},
	}
}

//AddItem ...
func (m *QueueManager) AddItem(item int, priority int) error {
	if priority > 4 {
		return errors.New("please you must sent a priority from 0 to 4")
	}

	m.PriorityQueues[priority].Enqueue(item)
	return nil
}

//GetHigherPriorityItem ...
func (m *QueueManager) GetHigherPriorityItem() *int {
	for i := range m.PriorityQueues {
		if !m.PriorityQueues[i].IsEmpty() {
			item := m.PriorityQueues[i].Dequeue()
			return item
		}
	}
	return nil
}

func main() {
	manager := NewQueueManager()
	manager.AddItem(40, 4)
	manager.AddItem(30, 3)
	manager.AddItem(20, 2)
	manager.AddItem(999, 0)
	manager.AddItem(10, 1)
	manager.AddItem(17, 1)
	manager.AddItem(4, 0)

	item := manager.GetHigherPriorityItem()
	for item != nil {
		fmt.Println(*item)
		item = manager.GetHigherPriorityItem()
	}
}
