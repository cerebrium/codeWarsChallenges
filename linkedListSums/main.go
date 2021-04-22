package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// linked list structure
type Node struct {
	value int64
	next  *Node
}

// helper function for creating pointers to linked list nodes - returns pointer to node
func nodeBuilder(inputValue int64) *Node {

	// make the node
	node := new(Node)

	// add the input value to the node
	node.value = inputValue

	// return the node
	return node
}

// function for traversing the linked list
func sumNext(n1 *Node, n2 *Node) int64 {
	// returns the sum of the two nodes
	return n1.value + n2.value
}

// recursive function to return the
func returnLinkedSum(n1 *Node, n2 *Node, currentNode *Node, headNode *Node) *Node {
	// add the defer to allow for optimization
	defer wg.Done()

	// add one more the the wait group in any of these cases
	wg.Add(1)

	// case to break the recursion
	if n1.next == nil && n2.next == nil {
		if sumNext(n1, n2) > 9 {
			// increment the current node to account for the ten
			currentNode.value += 1

			// add the new value onto the list
			currentNode.next = nodeBuilder(sumNext(n1, n2) - 10)

			// return the final list
			return headNode
		} else {
			// add the final node
			currentNode.next = nodeBuilder(sumNext(n1, n2))

			// return the final list
			return headNode
		}
	}

	// case if node one is nil and the other isnt
	if n1 == nil && n2 != nil {
		// add the current value to the list
		currentNode.next = nodeBuilder(n2.value)

		// return the function
		return returnLinkedSum(nil, n2.next, currentNode.next, headNode)
	}

	// case if node two is nil and the other isnt
	if n1 != nil && n2 == nil {
		// add the current value to the list
		currentNode.next = nodeBuilder(n1.value)

		// return the function
		return returnLinkedSum(n1.next, nil, currentNode.next, headNode)
	}

	if sumNext(n1, n2) > 9 {
		// increment the current node to account for the ten
		currentNode.value += 1

		// add the new value onto the list
		currentNode.next = nodeBuilder(sumNext(n1, n2) - 10)

		// return the function
		return returnLinkedSum(n1.next, n2.next, currentNode.next, headNode)
	} else {
		// if sum is lower than nine add and call
		currentNode.next = nodeBuilder(sumNext(n1, n2))

		// return the function
		return returnLinkedSum(n1.next, n2.next, currentNode.next, headNode)
	}
}

// function to display the number
func createTheResult(list *Node, resultString string) string {

	// this can also be optimized
	defer wg.Done()

	// only this case is recursive
	wg.Add(1)

	if list.next != nil {

		// add the current value to the string
		resultString += strconv.Itoa(int(list.value))

		// return the string builder
		return createTheResult(list.next, resultString)
	} else {
		// add the current value to the string
		resultString += strconv.Itoa(int(list.value))

		// return the string
		return resultString
	}

}

// make a wait group
var wg sync.WaitGroup

func main() {
	// tells the processor to use all the available cpu's
	runtime.GOMAXPROCS(runtime.NumCPU())

	// make a list for 9396247
	listOne := nodeBuilder(9)
	listOne.next = nodeBuilder(3)
	listOne.next.next = nodeBuilder(9)
	listOne.next.next.next = nodeBuilder(6)
	listOne.next.next.next.next = nodeBuilder(2)
	listOne.next.next.next.next.next = nodeBuilder(4)
	listOne.next.next.next.next.next.next = nodeBuilder(7)

	// make a second list 5427828
	listTwo := nodeBuilder(5)
	listTwo.next = nodeBuilder(4)
	listTwo.next.next = nodeBuilder(2)
	listTwo.next.next.next = nodeBuilder(7)
	listTwo.next.next.next.next = nodeBuilder(9)
	listTwo.next.next.next.next.next = nodeBuilder(2)
	listTwo.next.next.next.next.next.next = nodeBuilder(8)

	// make new head for final list
	resultList := nodeBuilder(0)

	// since the function counts as one this is defered as well
	defer wg.Done()

	// add one for initial call
	wg.Add(1)

	// call the functions to return the list
	fmt.Println("sum: ", createTheResult(returnLinkedSum(listOne, listTwo, resultList, resultList), ""))
}
