package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("> What is the quote?")
	fmt.Print("< ")
	quote, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("> Who said it?")
	fmt.Print("< ")
	performer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("> %v says, %q.\n", strings.TrimSpace(performer), strings.TrimSpace(quote))

}
