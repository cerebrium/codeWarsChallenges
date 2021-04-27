package main

import "fmt"

// question: given an array, and a number, find the two numbers in the array that come together to make the number given

// Thoughts:

// 1. this is acheivable via a double loop, but would be extremely inneficient
// 	a. alternatives
// 		1. make a hashtable
// 			a. given the sum find in the hashtable if the needed counterpart is available
// 		2. use binary search to find the needed number

type hashMapType struct {
	present bool
	index   int
}

// make a pure function that takes an array and returns a hashtable
func hashConvert(numList []int32) map[int32]hashMapType {
	// instantiate the hash to return
	returnHash := make(map[int32]hashMapType)

	// loop through the array and make the values in the hash
	for index, ele := range numList {
		returnHash[ele] = hashMapType{
			present: true,
			index:   index,
		}
	}

	// return the map
	return returnHash
}

// make a function that accepts an array and number
func twoSum(numList []int32, paramNum int32) []int {
	var finalList []int

	// make a hastable of the numbers available
	var intHash = hashConvert(numList)

	// take the values in the
	for index, element := range numList {

		// check if the numbers we want are present
		if intHash[paramNum-element].present {

			// if we find them, add the two indices to the array
			finalList = append(finalList, index, intHash[paramNum-element].index)

			// return the array
			return finalList
		}
	}

	// return the empty array if they aren't there
	return finalList
}

func main() {
	// should return [1, 2]
	practiceArray := []int32{2, 7, 11, 15}

	// call the function
	fmt.Println(twoSum(practiceArray, 18))

}
