package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type iris struct {
	Slength string
	Swidth  string
	Plength string
	Pwidth  string
	Variety string
}

func main() {
	file, err := os.Open("iris.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comment = '#'
	r.Comma = ','

	var list []iris
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("Bad Column: ", pe.Column)
				fmt.Println("Bad Line: ", pe.Line)
				fmt.Println("Error reported: ", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
			log.Fatal(err)
		}
		list = append(list, iris{record[0], record[1], record[2], record[3], record[4]})

	}

	fmt.Printf(" > First struct is : %v\n > Last struct is  : %v\n > Slice lenght is : %v\n ", list[1], list[len(list)-1], len(list))

}
