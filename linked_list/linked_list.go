package linked_list

import "fmt"

// Node represents a node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head   *Node
	length int
}

// prints the values in the linked list
// bt value of the node
func (l LinkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toPrint.Data)
		toPrint = toPrint.Next
		l.length--
	}

	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {
	
	// if the list is empty
	if l.length == 0 {
		return
	}

	if l.head.Data == value {
		l.head = l.head.Next // delete the head
		l.length-- // decrement the length
		return 
	}

	// if the list is not empty	
	previousDeletedValue := l.head
	
	// traverse the list
	for previousDeletedValue.Next.Data != value {
		// if the value is not found
		if previousDeletedValue.Next.Next == nil {
			return
		}

		// if the value is found
		previousDeletedValue = previousDeletedValue.Next
	}

	// delete the value
	previousDeletedValue.Next = previousDeletedValue.Next.Next
	l.length--
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}
