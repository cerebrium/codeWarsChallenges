package main

import (
	"fmt"
	"sort"
)

/*

	Prompt: given an array of arrays, each inner array has two numbers, the first is their height,
	the second is the amount of people they can see in front of them. They can only see people as tall or
	taller than them

	steps:
		1. make a hashmap of the array
		2. insert the tallest first into the final array -> interate down to the shortest
		3. When inserting follow these rules:
			a. is the condition of the second number (amount of people as taller or taller in front) met?
				1. yes ? insert at index
				2. no ? continue and if number of people in front has been partially met by
				current index, decrease how many more indeices to pass

	First Solution: => To undergo optimization in further pushes

*/

// make a function that creates a map of sorted arrays
func hashGen(array [][]int32) map[int32][]int32 {

	// instantiate the map
	hash := map[int32][]int32{}

	// take values and map them to the map
	for _, ele := range array {

		// check if hash has the height yet
		_, present := hash[ele[0]]

		// if the element is found
		if present {

			// check if added or not
			added := false

			// add the value to the array -> needs to be in order
			for index, val := range hash[ele[0]] {

				// check each value and place it at index needed
				if val > ele[1] {

					// insert an element
					hash[ele[0]] = append(hash[ele[0]], 0)

					// shift elements
					copy(hash[ele[0]][index+1:], hash[ele[0]][index:])

					// place element where it belongs
					hash[ele[0]][index] = ele[1]

					// change the bool
					added = true

					// break the loop
					break
				}
			}

			if !added {

				// if the element hasnt already been added
				hash[ele[0]] = append(hash[ele[0]], ele[1])
			}

		} else {
			// if the index isn't found create an array
			innerArray := make([]int32, 0)
			innerArray = append(innerArray, ele[1])

			hash[ele[0]] = innerArray
		}
	}

	// return the hashmap of the array
	return hash
}

// reQue
func reQue(array [][]int32) [][]int32 {

	// instantiate array to reeturn
	var returnArray [][]int32

	// get the hashmap of the array
	var hashArray = hashGen(array)

	// make array of keys - for reversing the order
	keys := make([]int32, 0, len(hashArray))

	// make the key array
	for k := range hashArray {
		keys = append(keys, k)
	}

	// sort the keys
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })

	// loop through the key array from greatest -> least
	for i := 0; i < len(keys); i++ {
		fmt.Println("i: ", keys[i])

		// loop through the hash array
		for _, val := range hashArray[keys[i]] {
			// if it is the first element, insert it in
			if len(returnArray) == 0 {

				// make an array
				initialArray := []int32{keys[i], val}

				// add the tuple
				returnArray = append(returnArray, initialArray)
			} else {
				// handle the logic
				counter := val

				index := 0

				// loop through the current list
				for _, tuple := range returnArray {

					// if the counter is 0 break and add at index
					if counter == 0 {
						break
					}

					// iloop through and increment the counter each time the height is found
					if tuple[0] >= keys[i] {
						counter--
					}

					// increment the index
					index++
				}

				// add the tuple at the index
				inputArray := []int32{keys[i], val}

				// insert an element
				returnArray = append(returnArray, inputArray)

				// shift elements
				copy(returnArray[index+1:], returnArray[index:])

				// fmt.Println("array: ", inputArray, "index: ", index)

				// place element where it belongs
				returnArray[index] = inputArray
			}
		}
	}

	// retunr teh array
	return returnArray

}

func main() {
	// make the test array
	testArray := [][]int32{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}

	// call the function
	fmt.Println("solution: ", reQue(testArray))

	// expected output
	// [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

}
