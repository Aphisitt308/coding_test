package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	atoz := []string{"A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X",
		"Y", "Z"}

	// input data
	fmt.Printf("Enter your string : ")
	fmt.Scanln(&input)
	Str := strings.ToUpper(input)
	fmt.Printf("\n")

	for i := range atoz {
		count := strings.Count(Str, atoz[i])
		if count != 0 {
			fmt.Println(atoz[i], count)
		}
	}

}
