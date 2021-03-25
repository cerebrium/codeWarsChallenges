package main

import "fmt"

type Bank struct {
	TwentyFive int
	Fifty      int
	OneHundred int
	Sum        int
}

func Tickets(peopleInLine []int) string {
	var myBank Bank

	for _, person := range peopleInLine {
		// handle base case
		if person == 25 {
			myBank.TwentyFive += 1
			myBank.Sum += 25
		} else {
			// handle all other cases
			if myBank.Sum > (myBank.Sum - person) {
				// handle the cases

				// 50
				if person == 50 {
					myBank.Fifty += 1
					if myBank.TwentyFive > 0 {
						myBank.TwentyFive -= 1
						myBank.Sum -= 25
					} else {
						return "NO"
					}
				}

				// 100
				if person == 100 {
					myBank.OneHundred += 1
					if myBank.TwentyFive > 0 {
						if myBank.Fifty > 0 {
							myBank.Fifty -= 1
							myBank.TwentyFive -= 1
							myBank.Sum -= 75
						} else {
							if myBank.TwentyFive > 2 {
								myBank.TwentyFive -= 3
								myBank.Sum -= 75
							} else {
								return "NO"
							}
						}
					} else {
						return "NO"
					}
				}
			} else {
				// not enough change
				return "NO"
			}
		}
	}
	return "YES"
}

func main() {
	fmt.Println(Tickets([]int{25, 25, 50}))
}
