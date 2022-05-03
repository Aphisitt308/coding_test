package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input int64
	fmt.Printf("Enter your number : ")
	fmt.Scan(&input)
	bistr := strconv.FormatInt(input, 2)
	strs := strings.Split(bistr, "")
	binary := make([]int, len(strs))
	for i := range binary {
		binary[i], _ = strconv.Atoi(strs[i])
		if binary[i] == 0 {
			binary[i] = 1
		} else {
			binary[i] = 0
		}
	}

	bistr1 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(binary)), ""), "[]")
	dec, _ := strconv.ParseInt(bistr1, 2, 64)
	fmt.Printf("\n")
	fmt.Println(dec)
}
