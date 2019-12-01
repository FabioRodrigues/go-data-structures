package main

import "fmt"

type vertex struct {
	distance int
	children []string
}

var graph = map[string]vertex{
	"A": vertex{distance: 7, children: []string{"B", "C"}},
	"B": vertex{distance: 3, children: []string{"D", "E"}},
	"C": vertex{distance: 12, children: []string{"F", "G"}},
	"D": vertex{distance: 99, children: []string{}},
	"E": vertex{distance: 8, children: []string{}},
	"F": vertex{distance: 5, children: []string{}},
	"G": vertex{distance: 6, children: []string{}},
}

func (v vertex) GetMaximumPath() {
	if len(v.children) == 0 {
		return
	}

	var maxVertex *vertex

	for _, c := range v.children {
		vertex := graph[c]
		if maxVertex == nil {
			maxVertex = &vertex
			continue
		}

		if maxVertex.distance > vertex.distance {
			maxVertex = &vertex
		}
	}

	fmt.Println("I'm going through path", maxVertex.distance)
	maxVertex.GetMaximumPath()

}

func main() {
	root := graph["A"]
	root.GetMaximumPath()
}
