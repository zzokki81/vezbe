package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("> What is the quote?")
	fmt.Print("< ")
	answer1, _ := reader.ReadString('\n')

	fmt.Println("> Who said it?")
	fmt.Print("< ")
	answer2, _ := reader.ReadString('\n')

	fmt.Printf("> %v says, %s", answer2[:len(answer2)-1], answer1)

}
