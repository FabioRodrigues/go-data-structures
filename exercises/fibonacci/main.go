package main

import "fmt"

func main() {
	fib := fibonacciRecursive(9)
	fmt.Println(fib)
}

var isInCache map[int]bool = map[int]bool{}
var cache map[int]int = map[int]int{}

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
	if isInCache[term] {
		fmt.Printf("item %d is in cache\n", term)
		return cache[term]
	}
	fmt.Printf("item %d is not in cache\n", term)
	if term == 0 || term == 1 {
		isInCache[term] = true
		cache[term] = term
		return term
	}
	result := fibonacciRecursive(term-1) + fibonacciRecursive(term-2)
	isInCache[term] = true
	cache[term] = result

	return result
}
