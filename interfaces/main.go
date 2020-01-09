package main

import "fmt"

type magazine struct {
	name     string
	edition  string
	authors  []string
	articles map[uint8]string
}

type sector struct {
	category string
	magazine
}

func (s *sector) write() {
	fmt.Printf("  > Best magazine for 2019  in category %v is: %v %q\n  > From the authors %v and %v\n"+
		"  > And best articles is %v\n ", s.category, s.magazine.name, s.magazine.edition, s.magazine.authors[0], s.magazine.authors[1],
		s.magazine.articles)
}

type writer interface {
	write()
}

func passingInterface(w writer) {
	w.write()
}

func main() {
	magazineValues := &sector{
		"Fitnes&Helth",
		magazine{
			name:     "Men's Helth",
			edition:  "Platinum edition",
			authors:  []string{"Matt Goulding", "Ray Klerck"},
			articles: map[uint8]string{1: "Helth Food", 2: "Back in Shape"}},
	}
	passingInterface(magazineValues)
}
