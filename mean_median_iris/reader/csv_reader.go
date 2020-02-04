package reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/zzokki81/vezbe/mean_median_iris/model/iris"
)

func CsvReader(filename string) []*iris.Iris {

	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comment, r.Comma = '#', ','

	list := []*iris.Iris{}

	var listFloat []float64
	var listString []string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 4; i++ {
			num, err := strconv.ParseFloat(record[i], 64)

			if err != nil {
				log.Fatal(err)
			}
			listFloat = append(listFloat, num)
		}
		listString = append(listString, record[4])
	}

	for g, h := 0, 0; g < len(listFloat); g, h = g+4, h+1 {
		iris := iris.New(listFloat[g], listFloat[g+1], listFloat[g+2], listFloat[g+3], listString[h])
		list = append(list, iris)
	}

	return list
}
