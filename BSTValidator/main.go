package main

import (
	"fmt"
	"math"
)

/*
	steps make the node structure
	solve the bst issue

*/

// node structure
type Node struct {
	value int64
	left  *Node
	right *Node
}

// function for helping to create nodes
func nodeCreator(inputValue int64) *Node {

	// instatiate the node
	node := new(Node)

	// add the value
	node.value = inputValue

	// return a pointer to a node
	return node
}

// recursively call itself to check full tree for validity
func bstHelper(n *Node, low int64, high int64) bool {

	// catch the nil case since empty node is valid
	if n == nil {
		return true
	}

	// took me awhile to figure out returning the whole statement
	return n.value > low && n.value < high && bstHelper(n.left, low, n.value) && bstHelper(n.right, n.value, high)

}

func bstChecker(n *Node) bool {

	// return the recursive function
	return bstHelper(n, math.MinInt64, math.MaxInt64)
}

func main() {

	// false case
	Node := nodeCreator(5)
	Node.left = nodeCreator(4)
	Node.right = nodeCreator(3)

	// true case
	TrueNode := nodeCreator(5)
	TrueNode.left = nodeCreator(3)
	TrueNode.right = nodeCreator(7)

	// call them
	fmt.Println("false case: ", bstChecker(Node))
	fmt.Println("true case: ", bstChecker(TrueNode))
}
