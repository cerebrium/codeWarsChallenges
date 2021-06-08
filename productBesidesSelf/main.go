package main

import (
	"fmt"
)

// prompt: given an array of numbers, return an array of all numbers besides the current index
// strategy: sum them all, return the sum divided by the index
// Time: O(n*2)
// space: one integar

// restriction givem: no using the / operator

type sums struct {
	forward  map[int]int
	backward map[int]int
}

// helper function to sum the array
func findProduct(array []int) int {
	// initialize product as 1
	product := 1

	// loop through and find product of all the numbers
	for _, val := range array {
		product *= val
	}

	return product
}

// function using the / operator
// function to return the new array
func productSums(array []int) []int {
	// find the product sum of the array
	productSum := findProduct(array)

	// loop through and replace the values with the value they should be
	for index, value := range array {
		array[index] = productSum / value
	}

	return array
}

// hash map of pre-proccessed values
func hashMapper(array []int) sums {

	// instantiate the map
	hashForwards := make(map[int]int)

	// instantiate the map
	hashBackwards := make(map[int]int)

	// current value
	forwardValue := 1

	backwardValue := 1

	// loop through array
	for i, val := range array {
		// set forward value
		forwardValue *= val

		// set backward value
		backwardValue *= array[len(array)-(i+1)]

		// we arent interested in all of them summed
		if i != len(array)-1 {

			// set the first value
			hashForwards[i] = forwardValue

			// set other side
			hashBackwards[i] = backwardValue

		}

	}

	sumsMap := sums{
		forward:  hashForwards,
		backward: hashBackwards,
	}

	return sumsMap
}

// function to solve this without the / operator
// still want this to be in linear time
func productSumsRestricted(array []int) []int {

	finalArray := []int{}

	// create the hash
	hash := hashMapper(array)

	// make pointers because golang doesn't iterate through maps in order
	pointer := len(hash.backward) - 2
	frontPointer := 0

	// // loop hrough and produce array
	for key, value := range hash.forward {
		if key == len(hash.forward)-1 {
			finalArray = append(finalArray, value, hash.backward[key])
		} else {

			// increase the pointers only in this case
			finalArray = append(finalArray, hash.forward[frontPointer]*hash.backward[pointer])
			pointer -= 1
			frontPointer += 1
		}
	}

	return finalArray
}

func main() {
	// make sample array
	testArray := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	testArrayTwo := []int{6, 7, 8, 9}

	// call the function
	fmt.Println("expecting: [1814400 1209600 907200 725760 604800 518400 453600 403200 362880] ", productSums(testArray))
	fmt.Println("expecting: [504, 432, 378, 336] ", productSums(testArrayTwo))

	testArray = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	testArrayTwo = []int{6, 7, 8, 9}

	// without the / sighn
	fmt.Println("expecting: [1814400 1209600 907200 725760 604800 518400 453600 403200 362880] ", productSumsRestricted(testArray))
	fmt.Println("expecting: [504, 432, 378, 336] ", productSumsRestricted(testArrayTwo))
}
