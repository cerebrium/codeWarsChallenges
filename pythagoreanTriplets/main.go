package main

// Prompt: Pythagorean Triplets

// 	1. Initial solution: triple loop
// 		a. super slow - but space complexity is constant
// 	2. make an object with the squared array and check O(N^2)
// 		a. space-complexity: O(N) -> linear for the object
// 3. make a sorted squared array
// 	a. can then loop through once and should get linear for the final pass
// 	b. total time complexity is O(N^2*3)

// steps:
// 1. Make an array of squared numbers (sorted)
// 2. loop through and check the sums to the rest of the values

// function to sqaure the numbers in the array
func squareArray(array []int32) []int32 {
	// instantiate the array
	returnArray := []int32{}

	// loop through the array and square the number
	for _, val := range array {
		// square the number and sort it
		returnArray = sortArray(val*val, returnArray)
	}

	return returnArray
}

// function to sort the array
func sortArray(integer int32, array []int32) []int32 {
	// thoughts on time saving: check if the incoming number is closer to end or initial number in array and start from there

	// takes the squared number and places it in the array
	array = append(array, 0)   // Step 1
	copy(array[2:], array[1:]) // Step 2
	array[1] = integer

	return array
}

func main() {

}
