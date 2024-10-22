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
	if l.length == 0 {
		return
	}

	if l.head.Data == value {
		l.head = l.head.Next
		l.length--
		return
	}

	previousDeletedValue := l.head
	for previousDeletedValue.Next.Data != value {
		if previousDeletedValue.Next.Next == nil {
			return
		}
		previousDeletedValue = previousDeletedValue.Next
	}

	previousDeletedValue.Next = previousDeletedValue.Next.Next
	l.length--
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}
