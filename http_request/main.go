package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func comicFetch(url string) (*comic, error) {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var comicBook comic

	err = json.NewDecoder(response.Body).Decode(&comicBook)
	if err != nil {
		return nil, err
	}
	return &comicBook, nil
}

func comicLastFetch() (*comic, error) {

	response, err := comicFetch("https://xkcd.com/info.0.json")
	if err != nil {
		log.Fatal(err)
	}
	return response, nil
}

func main() {
	comicBook, err := comicLastFetch()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(comicBook)

}
