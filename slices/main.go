package main

import (
	"fmt"
	"vezbe/biggernumsl"
)

func main() {

	a := []int{1, 2, 16, 6, 19, 18, 20, 4, 5, 67}
	fmt.Println("Enter number :")
	var n int
	fmt.Scan(&n)
	b := biggernumsl.LargerNums(n, a)
	fmt.Println(b)
}
