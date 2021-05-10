package main

import "fmt"

/*
	find the non-unique in an array of duplicates
	Steps:
		1. make a hashmap
		2. find the unique

*/

// hashmap maker helper function
func hashMapper(array []int32) map[int32]bool {

	// declare hash
	hash := make(map[int32]bool)

	// loop through the array and map into map
	for _, ele := range array {
		// check if the hash has the value
		if hash[ele] {
			delete(hash, ele)
		} else {
			hash[ele] = true
		}
	}

	// return the hash
	return hash
}

// main function that checks the hashmap for the value
func nonDup(array []int32) int32 {
	// make a hashmap
	hash := hashMapper(array)

	// there is only one value in this hash
	for k := range hash {
		return k
	}

	// return nothing in case no duplicate
	return int32(-1)
}

func main() {
	testArray := []int32{4, 3, 2, 4, 1, 3, 2}

	fmt.Println("Answer: ", nonDup(testArray))
}

/*
	time-space complexity:
		the map deletes as it goes, so one loop is all it takes: O(N)

		the hashmap is all thats created, which as most would be same size as the array: O(N)

		time-space final: O(N)
*/
