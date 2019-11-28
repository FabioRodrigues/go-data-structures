package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
	fmt.Println(s.items)
}

func (s *Stack) pop() {
	if len(s.items) == 0 {
		return
	}

	s.items = s.items[:len(s.items)-1]
	fmt.Println(s.items)
}

func main() {
	s := Stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
}
