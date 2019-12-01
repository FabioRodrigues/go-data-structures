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

//Graph ...
type Graph struct {
	DimensionalRepresentation [][]int
}

//NewGraph ...
func NewGraph(representation [][]int) Graph {
	return Graph{
		DimensionalRepresentation: representation,
	}
}

//ExistsConnection ...
func (g Graph) ExistsConnection(x int, y int) bool {
	return g.DimensionalRepresentation[x][y] == 1
}

//GetDegreeOfAVertex ...
func (g Graph) GetDegreeOfAVertex(vertex int) int {
	if vertex > len(g.DimensionalRepresentation) {
		return 0
	}

	sum := 0
	for _, i := range g.DimensionalRepresentation[vertex] {
		sum += i
	}

	return sum
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

	//Example 1
	// matrix := [][]int{
	// 	{0, 1, 1, 1, 1, 1, 0, 0, 0, 0},
	// 	{0, 1, 1, 0, 1, 1, 0, 0, 0, 0},
	// }

	// graph := NewGraph(matrix)
	// fmt.Println(graph.GetDegreeOfAVertex(1))

	//Example 2 - breadth-first search
	myList := map[string][]person{
		"me":     []person{newPerson("alice", false), newPerson("bob", false), newPerson("claire", false)},
		"bob":    []person{newPerson("anuj", false), newPerson("peggy", false)},
		"alice":  []person{newPerson("peggy", false)},
		"claire": []person{newPerson("thom", false), newPerson("jonny", true)},
		"anuj":   []person{},
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

}
