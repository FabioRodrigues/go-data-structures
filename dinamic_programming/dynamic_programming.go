package main

import "fmt"

func main() {
	typedWord := []rune{'b', 'l', 'u', 'e'}
	suggestedWord := []rune{'c', 'l', 'u', 'e', 's'}

	table := [][]int{}

	//initializing the table with defailt zero values
	for i := 0; i < len(typedWord); i++ {
		line := []int{}
		for j := 0; j <= len(suggestedWord); j++ {
			line = append(line, 0)
		}
		table = append(table, line)
	}
	count := 0
	//checking if the letters intercepts
	for i := 0; i < len(typedWord); i++ {
		if typedWord[i] == suggestedWord[i] {
			if i == 0 {
				table[i][i] = 1
				count = 1
				continue
			}

			consequentValues := table[i-1][i-1] + 1
			if consequentValues > count {
				count = consequentValues
			}
			table[i][i] = consequentValues
			continue
		}
	}

	for _, e := range table {
		fmt.Println(e)
	}

	fmt.Println(count)

}
