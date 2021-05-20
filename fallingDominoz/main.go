package main

import "fmt"

// Prompt: Given a set of dominoz, if there are left anf right forces applied determine the outcome
// Data Structures: Array of letters: 'o' 'l' 'r'
// Solution one: loop with pointers at every force application
// go through each and apply logic

// structure for the force application
type Force struct {
	index     int
	direction string
	completed bool
}

// structure for handling tracking updates
type Update struct {
	index     int
	direction string
}

// main function
func fallingDominoz(array []string, forceArray []Force) []string {
	// keep track of time
	time := 1

	// apply time to the force
	// make an infinite loop to handle applying force while applicatble
	for time > 0 {

		// make a new map for potential changes
		updateMap := make(map[int]string)

		// make a count for ending the loop
		forcesCompleted := len(forceArray)

		// loop through the force inputs
		for _, force := range forceArray {

			// check if the completed value is true
			if !force.completed {

				// add the attempted update to update array --> LEFT
				if force.direction == "L" {

					// check for potential collision
					if val, ok := updateMap[force.index-(time-1)]; ok {

						// if there is the value
						if val == "R" {

							// should always be the opposite
							updateMap[force.index-(time-1)] = "O"
						}

						// change the two forces to be completed
						// get current force
						force.completed = true

						// find competing force
						for i, forceFound := range forceArray {

							if forceFound.index+(time-1) == force.index-(time-1) {
								forceArray[i].completed = true
							}
						}
					} else {
						// handle non-collision case
						updateMap[force.index-(time-1)] = "L"
					}
				}

				// add the attempted update to update array --> LEFT
				if force.direction == "R" {

					// check for potential collision
					if val, ok := updateMap[force.index+(time-1)]; ok {

						// if there is the value
						if val == "L" {

							// should always be the opposite
							updateMap[force.index+(time-1)] = "O"
						}

						// change the two forces to be completed
						// get current force
						force.completed = true

						// find competing force
						for i, forceFound := range forceArray {
							if forceFound.index-(time-1) == force.index+(time-1) {
								forceArray[i].completed = true
							}
						}
					} else {
						// handle non-collision case
						updateMap[force.index+(time-1)] = "R"
					}
				}
			} else {
				forcesCompleted -= 1
			}
		}

		// break the loop if no forces left
		if forcesCompleted == 0 {
			time = -1
		}

		// apply the changes
		for key, val := range updateMap {

			// edge detection for out of bounds errors
			if key > len(array)-1 || key < 0 {
				if key > len(array)-1 {
					// find competing force
					for i, forceFound := range forceArray {
						if forceFound.index+(time-1) == key {

							// make the cometing force no longer valid
							forceArray[i].completed = true
						}
					}
				} else {
					// find competing force
					for i, forceFound := range forceArray {
						if forceFound.index-(time-1) == key {

							// make the cometing force no longer valid
							forceArray[i].completed = true
						}
					}
				}

			} else {
				if array[key] == "O" {
					array[key] = val
				} else {
					// case collision with already turned domino
					// find competing force
					for i, forceFound := range forceArray {
						if forceFound.index-(time-1) == key || forceFound.index+(time-1) == key {

							// make the cometing force no longer valid
							forceArray[i].completed = true
						}
					}
				}
			}
		}

		// increment the loop
		time += 1
	}

	// return the updated array
	return array
}

// function to make forces
func forceMaker(index int, direction string, completed bool) Force {
	force := Force{
		index:     index,
		direction: direction,
		completed: completed,
	}

	return force
}

func main() {
	// make an array full of dominoz
	dominoArray := []string{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"}

	// make a force array
	forceArray := []Force{forceMaker(1, "R", false), forceMaker(7, "L", false), forceMaker(8, "R", false)}

	fmt.Println("falling dominoz: ", fallingDominoz(dominoArray, forceArray))
}

// Time-Space-Complexity:
// Time: Loop through the force array up to n times O(N*M^2)
// Space: make a hashmap O(N)
