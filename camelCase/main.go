package main

import (
	"fmt"
	"strings"
)

func ToCamelCase(s string) string {
	if strings.Contains(s, "-") {
		sEdit := strings.Split(s, "-")

		finalString := ""
		for index, element := range sEdit {
			if index != 0 {
				finalString = finalString + strings.Title(element)
			} else {
				// if string(element[0]) == strings.ToUpper(string(element[0])) {
				// 	finalString = finalString + strings.Title(element)
				// } else {
				// 	finalString = finalString +
				// }
				finalString = finalString + element
			}
		}
		return finalString
	} else {
		sEdit := strings.Split(s, "_")

		finalString := ""
		for index, element := range sEdit {
			if index != 0 {
				finalString = finalString + strings.Title(element)
			} else {
				finalString = finalString + element
			}
		}
		return finalString
	}
}

func main() {
	fmt.Println(ToCamelCase("The_stealth_warrior"))
	fmt.Println(ToCamelCase("The-stealth-warrior"))
}
