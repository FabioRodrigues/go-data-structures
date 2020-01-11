package main

import "fmt"

//5! => 5*4*3*2*1=120

func main() {
	fact := factorialNotRecursive(1)
	fmt.Println(fact)
}

func factorialNotRecursive(fact int) int {

	ret := fact
	for i := fact - 1; i > 0; i-- {
		ret *= i
	}
	return ret
}
