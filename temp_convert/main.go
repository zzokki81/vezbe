package main

import (
	"fmt"
	tf "vezbe/temp_convert/tempfunctions"
)

func main() {
	err := tf.Cases()
	if err != nil {
		panic(err)
	}
	fmt.Println("Thank you for using temperature converter :)")
}
