package linked_list

import "fmt"

type Node struct {
	Value any
	Next  *Node
}

func (n *Node) newNode(value any) *Node {
	return &Node{Value: value, Next: nil}
}

type LinkedList struct {
	current *Node
}

func (ll *LinkedList) Next() any {
	if ll.current == nil {
		fmt.Println("next is nil")
		return nil
	}
	value := ll.current.Value
	ll.current = ll.current.Next
	return value
}

func (ll *LinkedList) Push(value any) {
	if ll.current.Next == nil {
		ll.current.Next = ll.current.newNode(value)
	}

}

func (ll *LinkedList) HasNext() bool {
	return ll.current != nil
}

func List() {
	//create LL
	myList := LinkedList{&Node{1, &Node{2, &Node{3, nil}}}}
	for myList.HasNext() {

		fmt.Println("Element:", myList.current.Value)
		myList.Next()
	}

}
