package linked_list

import "fmt"

type Node struct {
	Value any
	Next  *Node
}

func NewNode(value any) *Node {
	return &Node{value, nil}
}

func (n *Node) PushLinkedList(val any) {
	if n.Next == nil {
		n.Next = NewNode(val)
		return
	}
	for {
		if n.Next == nil {
			n.Next = NewNode(val)
			return
		} else {
			n = n.Next
		}
	}
}

func (n *Node) Iterator() {
	for {
		if n.Next == nil {
			fmt.Println("Дошли до конца связного списка")
			return
		} else {
			fmt.Println("Получили:", n.Value)
			if n.Next.Next == nil {
				fmt.Println("Получили:", n.Next.Value)
			}
			n = n.Next
		}
	}
}

func (n *Node) Rec() {
	if n.Next == nil {
		return
	}
	fmt.Println("Value", n.Value)
	if n.Next.Next == nil {
		fmt.Println("Value", n.Next.Value)
	}
	n.Next.Rec()
}

//func (n *Node) PopLinkedList() *Node {}

func (n *Node) Print() {
	for {
		if n == nil {
			return
		}
		fmt.Println(n.Value)
		n = n.Next
	}
}
func LinkedList() {
	myList := NewNode(1)
	myList.PushLinkedList(2)
	myList.PushLinkedList(3)
	myList.PushLinkedList(4)
	//myList.Print()
	//myList.Iterator()
	myList.Rec()
	//fmt.Println(returnetNode.Value)
}
