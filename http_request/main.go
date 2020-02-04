package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Comic struct {
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

type xkcdClient struct {
	host string
}

const path = "info.0.json"

func fetchComic(url string) (*Comic, error) {

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

	var comic Comic

	if err = json.NewDecoder(response.Body).Decode(&comic); err != nil {
		return nil, err

	}
	return &comic, nil
}

func fetchLastComic() (*Comic, error) {
	adress := xkcdClient{"https://xkcd.com/"}

	response, err := fetchComic(adress.host + path)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func getComics(comic *Comic, slice *[]Comic) error {

	for i := comic.Num - 1; i > comic.Num-5; i-- {
		url := fmt.Sprintf("http://xkcd.com/%d/info.0.json", i)

		tempComic, err := fetchComic(url)
		if err != nil {
			log.Fatal(err)
			return err
		}
		*slice = append(*slice, *tempComic)
	}
	return nil
}

func main() {

	comicBook, err := fetchLastComic()
	if err != nil {
		log.Fatal(err)
	}
	comics := []Comic{}
	comics = append(comics, *comicBook)

	err1 := getComics(comicBook, &comics)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Last five:\n")
	for i, lastFive := range comics {
		fmt.Print(i+1, ": ")
		fmt.Println(lastFive)
	}
}
