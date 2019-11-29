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

func main() {
	root := NewNode(4)
	root.Insert(2)
	root.Insert(3)
	root.Insert(1)
	root.Insert(6)
	root.Insert(5)
	root.Insert(7)

	tree, err := json.Marshal(root)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(string(tree))

	fmt.Printf("\n.......... *** .......... \n\n\n")
	root.SearchKey(3)
}
