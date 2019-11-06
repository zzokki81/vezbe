package main

import (
	"fmt"
	"vezbe/biggernumsl"
)

func main() {

	a := []int{4, 6, 2, 8, 22, 45, 68, 74, 23}
	b := biggernumsl.LargerNums(5, a)

	fmt.Println(b)
}
