package main

import "fmt"

//Queue ...
type Queue struct {
	items []person
}

func newQueue() Queue {
	return Queue{
		items: []person{},
	}
}

//Enqueue ...
func (q *Queue) Enqueue(p person) {
	q.items = append(q.items, p)
}

//Dequeue ...
func (q *Queue) Dequeue() *person {
	if len(q.items) > 0 {
		item := q.items[0]
		q.items = q.items[1:]
		return &item
	}
	return nil
}

type person struct {
	name                       string
	knowsHowToCreateAnAirplane bool
}

func newPerson(name string, knows bool) person {
	return person{
		name:                       name,
		knowsHowToCreateAnAirplane: knows,
	}
}

func main() {
	//Example 2 - breadth-first search
	myList := map[string][]person{
		"me":     []person{newPerson("alice", false), newPerson("bob", false), newPerson("claire", false)},
		"bob":    []person{newPerson("anuj", false), newPerson("peggy", false)},
		"alice":  []person{newPerson("peggy", false)},
		"claire": []person{newPerson("thom", false), newPerson("jonny", true)},
		"anuj":   []person{newPerson("bob", false)},
		"peggy":  []person{},
		"tom":    []person{},
		"jonny":  []person{},
	}

	myInvestigationQueue := newQueue()

	myInvestigationQueue.Enqueue(newPerson("me", false))

	nextPerson := myInvestigationQueue.Dequeue()
	for nextPerson != nil {
		if !nextPerson.knowsHowToCreateAnAirplane {
			fmt.Printf("%s doesnt know. Adding his friends to the list\n", nextPerson.name)

			for _, n := range myList[nextPerson.name] {
				fmt.Printf("adding %s to the queue of investigation\n", n.name)
				myInvestigationQueue.Enqueue(n)
			}
			nextPerson = myInvestigationQueue.Dequeue()
			continue
		}

		fmt.Println("***********************************")
		fmt.Printf("\n%s knows how to create an airplane\n\n", nextPerson.name)
		fmt.Println("***********************************")
		break
	}
	fmt.Println("there's no one here whose is able to build an airplane")
}
