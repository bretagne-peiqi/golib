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

func Delete(t *Node, v int) bool {
	if t == nil {
		return false
	}

	parent := t
	found := false
	for {
		if t == nil {
			break
		}
		if v == t.Data {
			found = true
			break
		}

		parent = t
		if v < t.Data { //left
			t = t.Left
		} else {
			t = t.Right
		}
	}

	if found == false {
		return false
	}
	return deleteNode(parent, t)
}

func deleteNode(parent, t *Node) bool {
	if t.Left == nil && t.Right == nil {
		fmt.Println("delete node which has neither left nor right subtree")
		if parent.Left == t {
			parent.Left = nil
		} else if parent.Right == t {
			parent.Right = nil
		}
		t = nil
		return true
	}

	if t.Right == nil { // Right is vide
		fmt.Println("delete node which has only left subnodes ")
		parent.Left = t.Left
		//parent.Data = t.Left.Data
		//parent.Right = t.Left.Right
		t.Left = nil
		t = nil
		return true
	}

	if t.Left == nil { // Left is vide
		fmt.Println("delete node which has only right subnodes ")
		//parent.Left = t.Right.Left
		//parent.Data = t.Right.Data
		parent.Right = t.Right
		t.Right = nil
		t = nil
		return true
	}

	fmt.Println("delete node which has right and left subtree")
	previous := t
	next := t.Left
	for {
		if next.Right == nil {
			break
		}
		previous = next
		next = next.Right
	}

	t.Data = next.Data
	if previous.Left == next {
		previous.Left = next.Left
	} else {
		previous.Right = next.Right
	}
	next.Left = nil
	next.Right = nil
	next = nil
	return true
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
	var original = []int{44, 2, 3, 4, 7, 8, 1, 9, 99, 100, 5, 6}
	for _, v := range original {
		tree = Insert(tree, v)
	}
	Print(tree)
	ok := Search(tree, 2)
	fmt.Printf("search result is %v \n", ok)
	Delete(tree, 4)
	Print(tree)
	max, bo := GetMax(tree)
	min, boo := GetMin(tree)
	if bo == nil && boo == nil {
		fmt.Printf("max value and min value are %v and %v \n", max, min)
	}
}
