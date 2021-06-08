package main

import "fmt"

// prompt: given an array of numbers, return an array of all numbers besides the current index
// strategy: sum them all, return the sum divided by the index
// Time: O(n*2)
// space: one integar

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

func main() {
	// make sample array
	testArray := []int{2, 3, 4, 5}
	testArrayTwo := []int{6, 7, 8, 9}

	// call the function
	fmt.Println("expecting: [60, 40, 30, 24] ", productSums(testArray))
	fmt.Println("expecting: [504, 432, 378, 336] ", productSums(testArrayTwo))
}
