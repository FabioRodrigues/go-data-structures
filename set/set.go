package main

import "fmt"

//Collection ...
type Collection struct {
	Items map[int]bool
}

//NewCollection ..
func NewCollection() Collection {
	return Collection{
		Items: map[int]bool{},
	}
}

//Add ...
func (c *Collection) Add(i int) {
	c.Items[i] = true
}

//Join ...
func (c *Collection) Join(aditional []int) {
	for _, i := range aditional {
		c.Add(i)
	}
}

//Delete ...
func (c *Collection) Delete(i int) {
	delete(c.Items, i)
}

//Exists ...
func (c Collection) Exists(i int) bool {
	return c.Items[i]
}

//Size ...
func (c Collection) Size() int {
	return len(c.Items)
}

//IsEmpty ...
func (c Collection) IsEmpty() bool {
	return len(c.Items) == 0
}

func main() {
	coll := NewCollection()

	coll.Add(1)
	coll.Join([]int{2, 3, 4, 5, 1, 6, 7, 8, 8, 8})
	coll.Delete(8)
	fmt.Println(coll.Items)
	fmt.Printf("Is it empty? %v\nWhat's the size of it? %2d\n", coll.IsEmpty(), coll.Size())
}
