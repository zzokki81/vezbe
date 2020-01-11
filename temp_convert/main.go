package main

import (
	"fmt"
	tf "github.com/zzokki81/vezbe/tempfunctions"
)

func main() {
	err := tf.Cases()
	if err != nil {
		panic(err)
	}
	fmt.Println("Thank you for using temperature converter :)")
}
