package main

type Node struct {
	Data int
	Next *Node
}

func main2(head *Node) {

	var ptr, p, q *Node
	p = head.Next
	q = nil
	head.Next = nil

	for {
		if p != nil {
			ptr = p.Next
			p.next = q
			q = p
			p = ptr

		} else {
			break

		}
	}
	head.Next = q
}

func main(head *Node) {
	var ptr, p, q *Node
	p = head.Next
	q = nil
	head.Next = nil

	for {
		if p != nil {
			ptr = p.Next
			p.Next = q
			q = p
			p = ptr
		} else {
			break
		}
	}
	head.Next = q
}
