package main

import (
	"fmt"
	"sort"
)

func bSHelper(array []int32, number int32) int {
	// make index to help count the removed sections
	index := -1

	// find the middle of the array and check if number is higher or lower
	for len(array) > 1 {
		// if we found our index
		if array[len(array)/2] == number {
			// if there is a count missing add it on
			return index + len(array)/2
		}
		// if the number is less than the found number reduce the array to the left
		if number < array[len(array)/2] {
			array = array[0 : len(array)/2]
		}

		// if the number is greater than the found number reduce the array to the right
		if number > array[len(array)/2] {
			index += len(array) / 2
			array = array[len(array)/2:]
		}
	}

	return index
}

func binarySearch(array []int32, number int32) []int {
	// instantiate the index to return
	var indices []int

	// get the index of the first found case of the number
	var firstReturn = bSHelper(array, number)

	// if the number is found append it to the indice list
	if firstReturn > 0 {
		indices = append(indices, bSHelper(array, number))
	} else {

		// if the number isn't in the array return here
		return indices
	}

	// then find the end index
	for index, ele := range array[firstReturn:] {

		// iterate until not the current number
		if ele != number {

			// check if the number is in the array
			if index != 0 {

				// if non-number match found return the index of the previous match
				if len(indices) > 1 {
					indices[1] = firstReturn + index - 1
				} else {
					indices = append(indices, firstReturn+index-1)
				}
			} else {

				// there aren't two of the number in the array, return first index only
				return indices
			}
		}
	}

	// return the index
	return indices
}

func main() {
	// make annarray to work with
	array := []int32{1, 2, 3, 4, 3, 9, 4, 6, 7, 4, 5, 89, 9, 43, 1, 3, 5, 9, 8, 1, 3, 2, 9, 9, 9, 9, 9}

	// sort the array as part of the prompt
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })

	// call the returned function
	fmt.Println(binarySearch(array, 9))

}
