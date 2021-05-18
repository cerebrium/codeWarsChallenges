package main

import "fmt"

// Prompt:
//    Given an array [[1, 0]] where:
// 		the first number is course to take
// 		second number is the pre-requisites

// Determine a way to asses if there is a cycle [[1,0], [0,1]]

// {
// 	1: 0,
// 	0: 1
// }

// Thoughts
// can make a hashtable and check the requirments not being circular there

func hashMaker(array []int16, hMap map[int16][]int16) map[int16][]int16 {
	// add the value to the map
	if val, ok := hMap[array[0]]; ok {

		// if the key already exists, add the new requirment
		hMap[array[0]] = append(val, array[1])
	} else {
		// if the key does not exist, make requirment array
		newArray := []int16{array[1]}

		// adda the requirment array to map
		hMap[array[0]] = newArray
	}

	// return the modified map
	return hMap
}

// function to help determine if cycle exists
func checkCycle(class int16, preRequisite int16, hash map[int16][]int16) bool {
	// loop through prerequisite array to look for former prerequisite
	if checkVal, ok := hash[preRequisite]; ok {
		// loop through found array and if it contains the course return false
		for _, innerValue := range checkVal {
			// check for match
			if innerValue == class {

				// if match foud
				return true
			}
		}
	}

	// implicitly return false
	return false
}

// main function, takes the array of courses
func courseSchedule(array [][]int16) bool {
	// instantiate the hash to check from
	hash := make(map[int16][]int16)

	// loop through the matrix and add the elements to the hash
	for _, ele := range array {
		// call the hashing function on the course element
		hash = hashMaker(ele, hash)

		// call the cylce helper function
		if checkCycle(ele[0], ele[1], hash) {
			return false
		}
	}

	// implicitly return true
	return true
}

// function for determining text
func courseTextOutput(checkedSchedule bool) string {
	if checkedSchedule {
		return "Courses are viable"
	} else {
		return "Courses are circular"
	}
}

// helper function for making courses
func courseMaker(course int16, preRequisite int16) []int16 {
	// put the values into the array
	courseArray := []int16{course, preRequisite}

	// return array
	return courseArray
}

func main() {

	// make test array
	testArrayCircular := [][]int16{courseMaker(0, 1), courseMaker(1, 0)}
	testArrayViable := [][]int16{courseMaker(1, 0), courseMaker(2, 1)}

	// check for cycle
	fmt.Println("Should return circular: ", courseTextOutput(courseSchedule(testArrayCircular)))
	fmt.Println("Should return viable courses: ", courseTextOutput(courseSchedule(testArrayViable)))
}

// Space time complexity analysis
// 	Time: the worst case scenario is a loop through each element array, then array of matches O(n*m)
// 	Space: makes a hashmap of the array, so: O(n)
