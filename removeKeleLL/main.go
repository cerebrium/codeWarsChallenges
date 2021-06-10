package main

import (
	"fmt"
)

// Prompt: make a linked list where you can delete k element

// Thoughts:
//    1. Would like to do this with one pass O(n)
//    2. Iterate through and track index and pointers to two backward elements

// steps:
// 	  1. Create the data structure
// 	  2. Create helper function for the data structure creation
// 	  3. solve the problem

// time-space:
// Time: O(n)
// Space: creation of the linked list: O(n)

// -------------- Data Type Creation -------------------- //

type LNode struct {
	value int
	next  *LNode
}

// helper function for creating the nodes
func nodeCreator(value int) *LNode {

	// create a node
	node := new(LNode)

	// set the value to the node
	node.value = value

	// return the node
	return node
}

// takes an array and returns a linked list, returns head
func lCreater(array []int) *LNode {

	// keep a pointer to the last node, so can set its next
	pointer := nodeCreator(array[0])

	// create the head
	head := pointer

	// iterate through the list, create the rest of the linked list
	for i, val := range array {
		if i != 0 {
			// create a node from the int
			currentNode := nodeCreator(val)

			// point the previous node to the current node
			pointer.next = currentNode

			// change the pointer to be the current node
			pointer = currentNode
		}
	}

	// return the head
	return head

}

// function to help display the answer
func readLL(node *LNode) []int {
	// instantiate the slice to return
	array := []int{}

	// loop through the linked list and append to an array
	for node.next != nil {

		// add the value to the array
		array = append(array, node.value)

		// change the node
		node = node.next
	}

	// grab the last value
	array = append(array, node.value)

	// return the array
	return array
}

// ----------------- solving the question ------------------ //
func removeKEle(node *LNode, k int) *LNode {
	// keep the head
	head := node

	// handle edge case
	if k <= 1 {
		head = node.next
	} else {

		// keep two pointers
		twoBack := node
		node = node.next

		oneBack := node
		node = node.next

		counter := 3

		// iterate through linked list and find k
		for counter <= k && node.next != nil {

			// iterate the counter
			counter++

			// iterate the pointers
			twoBack = oneBack

			oneBack = node

			// call node.next
			node = node.next
		}

		// change the node to skip the k node
		twoBack.next = node

	}

	return head
}

func main() {

	// create a test array for ll creation
	testArray := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	// create the linked list
	linkedList := lCreater(testArray)

	// display the answer
	fmt.Println(readLL(removeKEle(linkedList, 1)))

}
