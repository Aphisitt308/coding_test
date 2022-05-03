package main

import "fmt"

func main() {
	var input uint
	fmt.Printf("Enter your money : ")
	fmt.Scan(&input)
	remain := input
	var sum uint
	banknote := []uint{100, 20, 10, 5, 1}
	fmt.Printf("\n")
	for i := range banknote {
		count := remain / banknote[i]
		remain = remain % banknote[i]
		sum += count
	}
	fmt.Println(sum)

}
