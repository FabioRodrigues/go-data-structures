package main

import "fmt"

//5! => 5*4*3*2*1=120

func main() {
	fact := factorialRecursive(5)
	fmt.Println(fact)
}

func factorialNotRecursive(fact int) int {

	ret := fact
	for i := fact - 1; i > 0; i-- {
		ret *= i
	}
	return ret
}

func factorialRecursive(fact int) int {
	if fact == 1 {
		return fact
	}
	return fact * factorialRecursive(fact-1)
}
