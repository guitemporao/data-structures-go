package main

import (
	"fmt"

	"github.com/guitemporao/go-dts/linked_list"
	"github.com/guitemporao/go-dts/queues"
	"github.com/guitemporao/go-dts/stacks"
)

func main() {

	// linked list
	myList := linked_list.LinkedList{}
	node1 := &linked_list.Node{Data: 1}
	node2 := &linked_list.Node{Data: 4}
	node3 := &linked_list.Node{Data: 5}

	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)

	myList.PrintListData()

	myList.DeleteWithValue(4)
	myList.PrintListData()

	// stacks
	newStack := stacks.Stack{}
	newStack.Push(1)
	newStack.Push(2)
	newStack.Push(3)
	fmt.Println(newStack.Pop()) // 3 -> remove the last index pushed in to the stack

	// queues
	newQueue := queues.Queue{}
	newQueue.Enqueue(1)
	newQueue.Enqueue(2)
	newQueue.Enqueue(3)
	newQueue.Enqueue(4)
	newQueue.Enqueue(5)
	newQueue.Enqueue(6)
	fmt.Println(newQueue.Dequeue()) // 1 -> remove the first index pushed in to the queue
	fmt.Println(newQueue.Peek())
}
