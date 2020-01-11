package main

import "fmt"

func main() {
	fib := fibonacciRecursive(9)
	fmt.Println(fib)
}

func fibonacciNotRecursive(term int) int {
	if term == 1 || term == 2 {
		return term - 1
	}
	terms := []int{1, 1}

	for i := 1; i < term-2; i++ {
		swap := terms[0] + terms[1]
		terms[0] = terms[1]
		terms[1] = swap
	}
	return terms[1]
}

func fibonacciRecursive(term int) int {
	if term == 0 || term == 1 {
		return term
	}
	return fibonacciRecursive(term-1) + fibonacciRecursive(term-2)
}
