package main

import "fmt"

// prompt: given an array of numbers, find the most frequent 'k' numbers (array is not necesarily sorted to start with)

// thoughts:
// A. can make a hashmap, rutern the highest count
// O(n*2) -> linear time, two passes (once for the array, once for the map)
// space: hashmap generated, so O(n)

// helper function to create this structure
func hashMapper(array []int) map[int]int {

	// the hashmap of all numbers
	hash := make(map[int]int)

	// loop through and create the hash
	for _, val := range array {

		// check if the value exist
		if value, ok := hash[val]; ok {
			hash[val] = value + 1
		} else {
			hash[val] = 1
		}
	}

	// return the struct with max hash
	return hash

}

// take the hash and return the k top values
func findKValues(array []int, amount int) []int {

	// make array to return
	returnArray := []int{}

	// make the hash
	hash := hashMapper(array)

	// make map of the highest
	finalMap := make(map[int]int)

	// loop through the hash and establish highest
	for key, value := range hash {
		// handle all k elements found
		if len(finalMap) == amount {

			// check the final map for lower values
			for fKey := range finalMap {
				if finalMap[fKey] < value {
					// add new value
					finalMap[key] = value

					// delete old value
					delete(finalMap, fKey)
				}
			}
		} else {
			finalMap[key] = value
		}
	}

	// make an array of the keys
	for key := range finalMap {
		returnArray = append(returnArray, key)
	}

	// return found top instances
	return returnArray
}

func main() {

	testArray := []int{1, 2, 1, 2, 3, 4, 5, 6, 7, 8, 1}

	fmt.Println("expecting: [1, 2] ", findKValues(testArray, 2))
}
