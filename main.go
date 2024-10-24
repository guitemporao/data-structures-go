package main

import (
	"fmt"
	"time"

	"github.com/guitemporao/go-dts/internal/currency"
)


func runConcurrencyWork(
	workerID int,  // worker id
	currencyChan <- chan currency.Currenncy, // read from channel
	resultChan chan <- currency.Currenncy, // write to channel
	) {
	
   // write to channel
   fmt.Printf("worker %d started\n", workerID)

   // read from channel
   for c := range currencyChan {
	   // fetch rates
	   rates, err := currency.FetchCurrencyRates(c.Code)
	   
	   // write to channel
	   if err != nil {
		   fmt.Printf("worker %d error: %v\n", workerID, err)
	   } else {
		   c.Rates = rates // set rates
		   resultChan <- c // write to channel
	   }

	   // write to channel
	   fmt.Printf("worker %d done\n", workerID)
   }
}


func main(){
	// create a currency struct
	ce := &currency.MyCurrency{
		Currencies: make(map[string]currency.Currenncy),
	}

	err := ce.FetchAllConcurrencies()

	if err != nil {
		fmt.Println(err)
		return
	}
	
	// create channels
	currencyChan := make(chan currency.Currenncy, len(ce.Currencies))
	resultChan := make(chan currency.Currenncy, len(ce.Currencies))

	// create workers
	for i := 0; i < len(ce.Currencies); i++ {
		go runConcurrencyWork(i, currencyChan, resultChan)
	}

	// send jobs
	startTime := time.Now()

	// wait for results
	resultTitme := 0; 

	// send jobs
	for _, curr := range ce.Currencies {
		currencyChan <- curr
		// resultTitme++
	}

	// wait for results
	for  {
		if resultTitme == len(ce.Currencies) {
				fmt.Println("closing resultChan")
				close(currencyChan)
				break
			}

		select {
			case c := <- resultChan:
				// time.Sleep(1 * time.Second)
				ce.Currencies[c.Code] = c
				resultTitme++

			case <- time.After(5 * time.Second):
				// time.Sleep(1 * time.Second)
				fmt.Printf("time out: %v\n", time.Since(startTime))
				return
		}
	}

	// close(resultChan)
	endTime := time.Now()

	fmt.Println("======== Results ========")

	// print results
	for _, curr := range ce.Currencies {
				fmt.Printf("%s (%s): %d rates\n", curr.Name, curr.Code, len(curr.Rates))
	}

	fmt.Println("=========================")
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}



// func main() {

// 	// linked list
// 	myList := linked_list.LinkedList{}
// 	node1 := &linked_list.Node{Data: 1}
// 	node2 := &linked_list.Node{Data: 4}
// 	node3 := &linked_list.Node{Data: 5}

// 	myList.Prepend(node1)
// 	myList.Prepend(node2)
// 	myList.Prepend(node3)

// 	myList.PrintListData()

// 	myList.DeleteWithValue(4)
// 	myList.PrintListData()

// 	// stacks
// 	newStack := stacks.Stack{}
// 	newStack.Push(1)
// 	newStack.Push(2)
// 	newStack.Push(3)
// 	fmt.Println(newStack.Pop()) // 3 -> remove the last index pushed in to the stack

// 	// queues
// 	newQueue := queues.Queue{}
// 	newQueue.Enqueue(1)
// 	newQueue.Enqueue(2)
// 	newQueue.Enqueue(3)
// 	newQueue.Enqueue(4)
// 	newQueue.Enqueue(5)
// 	newQueue.Enqueue(6)
// 	fmt.Println(newQueue.Dequeue()) // 1 -> remove the first index pushed in to the queue
// 	fmt.Println(newQueue.Peek())

// 	// binary search tree
// 	newNode := binary_search_tree.Node{Key: 100}
// 	newNode.Insert(105)
// 	newNode.Insert(45)
// 	fmt.Println(newNode)
// 	search := newNode.Search(45)
// 	fmt.Println(search)
// }
