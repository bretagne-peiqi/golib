package main

import (
	"fmt"
	//	"math/rand"
	//	"strconv"
)

type Node struct {
	data int
	next *Node
}

func quiksort(arra []int, left int, right int) []int {

	begin := arra[left]
	i, j := left+1, right

	for i < j {

		if arra[i] > begin {
			arra[i], arra[j] = arra[j], arra[i]
			j--
		} else {
			i++
		}
	}

	if arra[i] >= begin {
		i--
	}
	arra[i], arra[left] = begin, arra[i]

	if i-left > 1 {
		quiksort(arra, left, i-1)
	}
	if right-i > 1 {
		quiksort(arra, i+1, right)
	}
	return arra
}

// change value
func vAdapter(head *Node) {
	var arra []int

	for i := head; i != nil; i = i.next {
		arra = append(arra, i.data)
	}

	for _, x := range arra {
		fmt.Printf("orignal order is %v \n", x)
	}
	fmt.Printf("processing ... \n")
	ordered := quiksort(arra, 0, len(arra)-1)
	for _, w := range ordered {
		fmt.Printf("reorder order is %v \n", w)
	}
	j := head
	for _, item := range ordered {
		if j != nil {
			j.data = item
			j = j.next
		}

	}
}

func addNode(head *Node, seed int) {

	node := new(Node)
	node.next = nil
	node.data = seed

	tmptr := head
	for {
		if tmptr.next != nil {
			tmptr = tmptr.next
		} else {
			break
		}
	}
	tmptr.next = node
}

func main() {
	head := new(Node)
	head.next = nil
	head.data = 0
	orignal := []int{4, 5, 3, 8, 11, 2, 31, 9}
	for _, x := range orignal {
		//seed := rand.Int()
		addNode(head, x)
	}

	tmp := head
	ct := 0
	for {
		if tmp.next != nil {
			fmt.Printf("this is the %d and the value is %v \n", ct, tmp.data)
			ct++
			tmp = tmp.next
		} else {
			break
		}
	}
	fmt.Printf("this is the %d and the value is %v \n", ct, tmp.data)
	vAdapter(head)
	tmp2 := head
	ct = 0
	for {
		if tmp2.next != nil {
			fmt.Printf("this is the %d and the value is %v \n", ct, tmp2.data)
			ct++
			tmp2 = tmp2.next
		} else {
			break
		}
	}
	fmt.Printf("this is the %d and the value is %v \n", ct, tmp2.data)
}
