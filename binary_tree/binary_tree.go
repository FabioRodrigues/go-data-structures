package main

import (
	"encoding/json"
	"fmt"
)

//Node ...
type Node struct {
	Left  *Node
	Right *Node
	Key   int
}

//NewNode ...
func NewNode(key int) *Node {
	return &Node{Key: key}
}

//Insert ...
func (n *Node) Insert(key int) {
	fmt.Printf("inserting key %2d ...\n", key)
	//time.Sleep(5 * time.Second)

	if key < n.Key {
		fmt.Printf("key %2d is lower than %2d. Going Left\n", key, n.Key)
		if n.Left == nil {
			fmt.Printf("adding key %2d under %2d node (left side)\n", key, n.Key)
			n.Left = NewNode(key)
			return
		}
		n.Left.Insert(key)
		return
	}

	fmt.Printf("key %2d is greather than %2d. Going right\n", key, n.Key)
	if n.Right == nil {
		fmt.Printf("adding key %2d under %2d node (right side)\n", key, n.Key)
		n.Right = NewNode(key)
		return
	}
	n.Right.Insert(key)
	return
}

//SearchKey ...
func (n Node) SearchKey(key int) {
	//base case
	if n.Key == key {
		fmt.Printf("key %2d found!\n", key)
		return
	}

	if key < n.Key {
		fmt.Println("going left")
		if n.Left == nil {
			fmt.Println("key not found :(")
			return
		}
		n.Left.SearchKey(key)
		return
	}

	fmt.Println("going right")
	if n.Right == nil {
		fmt.Println("key not found :(")
		return
	}
	n.Right.SearchKey(key)
	return
}

//Delete ...
func (n *Node) Delete(key int) {
	if key < n.Key {
		if n.Left == nil {
			fmt.Printf("key %2d not found", key)
			return
		}
		if n.Left.Key == key {
			fmt.Println("deleting item", key)
			n.Left = nil
			return
		}
		n.Left.Delete(key)
		return
	}

	if n.Right == nil {
		fmt.Printf("key %2d not found", key)
		return
	}
	if n.Right.Key == key {
		fmt.Println("deleting item", key)
		n.Right = nil
		return
	}
	n.Right.Delete(key)
	return
}

func main() {

	root := insertItems([]int{4, 2, 3, 1, 6, 5, 7})
	// root := insertItems([]int{4, 2, 5})

	printItems(root)

	// root.SearchKey(3)
	root.Delete(3)
	printItems(root)

}

func printItems(root *Node) {
	tree, err := json.Marshal(root)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("\n.......... *** .......... \n\n\n")
	fmt.Println(string(tree))
	fmt.Printf("\n.......... *** .......... \n\n\n")
}

func insertItems(items []int) *Node {
	if len(items) == 0 {
		return nil
	}

	root := NewNode(items[0])

	for _, e := range items[1:] {
		root.Insert(e)
	}
	return root
}
