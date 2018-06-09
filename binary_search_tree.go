package main

import "fmt"

// Left node value is litter than root value, Right node value is always than root value
type Node struct {
	Left  *Node
	Data  int
	Right *Node
}

func Insert(tree *Node, Data int) *Node {
	if tree == nil {
		return &Node{nil, Data, nil}
	}
	if Data >= tree.Data {
		tree.Right = Insert(tree.Right, Data)
		return tree
	}
	tree.Left = Insert(tree.Left, Data)
	return tree
}

func Search(tree *Node, v int) bool {
	if tree == nil {
		return false
	}
	switch {

	case v > tree.Data:
		{
			return Search(tree.Right, v)
		}
	case v < tree.Data:
		{
			return Search(tree.Left, v)
		}
	default:
		return true
	}
	return false

}

func Delete(tree *Node, v int) bool {
	if tree == nil {
		return false
	}
	return false
}

func Print(tree *Node) {
	if tree == nil {
		return
	}
	Print(tree.Left)
	fmt.Printf("the value is %v \n", tree.Data)
	Print(tree.Right)
}

func GetMin(tree *Node) (int, error) {
	if tree == nil {
		return -1, fmt.Errorf("doest not exsit %v", tree)
	}
	for {
		if tree.Right != nil {
			tree = tree.Right
		} else {
			break
		}
	}
	return tree.Data, nil
}

func GetMax(tree *Node) (int, error) {
	if tree == nil {
		return -1, fmt.Errorf("doesn't exsit")
	}
	for {
		if tree.Left != nil {
			tree = tree.Left
		} else {
			break
		}
	}
	return tree.Data, nil
}

func main() {
	var tree *Node = nil
	var original = []int{44, 2, 3, 4, 7, 8, 1, 9, 99, 100}
	for _, v := range original {
		tree = Insert(tree, v)
	}
	Print(tree)
	ok := Search(tree, 2)
	fmt.Printf("search result is %v \n", ok)

	max, bo := GetMax(tree)
	min, boo := GetMin(tree)
	if bo == nil && boo == nil {
		fmt.Printf("max value and min value are %v and %v \n", max, min)
	}
}
