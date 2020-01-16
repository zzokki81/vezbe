package functions

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func CsvReader(filename string) [][]string {

	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comment, r.Comma = '#', ','

	list, counter := [][]string{}, 0

	for {
		record, err := r.Read()
		counter++
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
			log.Print(err)
		}
		list = append(list, record)
	}
	return list
}