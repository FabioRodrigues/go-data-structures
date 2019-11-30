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

func (n Node) hasJustOneChild() (bool, *Node) {

	if n.Right == nil && n.Left != nil {
		return true, n.Left
	}

	if n.Left == nil && n.Right != nil {
		return true, n.Right
	}
	return false, nil
}

func (n Node) hasnoChildren() bool {
	return (n.Right == nil && n.Left == nil)
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

//Min ...
func (n *Node) Min() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.Min()
}

//Max ...
func (n *Node) Max() *Node {
	if n.Right == nil {
		return n
	}
	return n.Right.Min()
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
			nodeToBeDeleted := *n.Left
			fmt.Println("deleting item", key)

			//first case, trying to remove a leaf
			if nodeToBeDeleted.hasnoChildren() {
				n.Left = nil
				return
			}

			//second case, has just one child
			has, child := nodeToBeDeleted.hasJustOneChild()
			if has {
				n.Left = child
				return
			}

			//last case, get the leftiest node of the right subtree
			min := *nodeToBeDeleted.Right.Min()
			n.Delete(min.Key)

			n.Left = &min
			n.Left.Left = nodeToBeDeleted.Left
			n.Left.Right = nodeToBeDeleted.Right

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

		nodeToBeDeleted := n.Right

		if n.Right.hasnoChildren() {
			n.Right = nil
			return
		}

		has, child := n.Right.hasJustOneChild()
		if has {
			n.Right = child
			return
		}
		//last case, get the leftiest node of the right subtree
		min := *nodeToBeDeleted.Right.Min()
		n.Delete(min.Key)

		n.Right = &min
		n.Right.Left = nodeToBeDeleted.Left
		n.Right.Right = nodeToBeDeleted.Right
		return
	}
	n.Right.Delete(key)
	return
}

func main() {

	root := insertItems([]int{50, 30, 100, 20, 40, 35, 45, 37})

	printItems(root)
	root.Delete(30)
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
