package main

import "fmt"

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

func main() {

	//Example 1
	matrix := [][]int{
		{0, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 1, 1, 0, 1, 1, 0, 0, 0, 0},
	}

	graph := NewGraph(matrix)
	fmt.Println(graph.GetDegreeOfAVertex(1))

}
