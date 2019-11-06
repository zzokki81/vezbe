package main

import "fmt"

func main() {

	unsort := []int{23, 11, 43, 52, 22, 67, 56, 34, 18}
	BubbleSort(unsort)
	fmt.Println("Before sorting : ", unsort)
	fmt.Println("After sorting : ", unsort)

}

type Sortable interface {
	Less(i, j int) bool
	Swap(i, j int)
}

type nums = []int 

func(n *num) Less(i, j int) {
	return n[i] < n[j]
}

 

func BubbleSort(unsort []int) {

	for {
		n := len(unsort) - 1

		swapped := false

		for i := 0; i < n; i++ {

			if unsort[i] > unsort[i+1] {

				swapped = true
				unsort[i], unsort[i+1] = unsort[i+1], unsort[i]
			}
		}

		if !swapped {
			break
		}
	}

}
