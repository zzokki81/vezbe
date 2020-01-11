package main

import "fmt"

func main() {

	a := []int{1, 3, 4, 6, 8, 16, 17, 31, 33, 39}
	target := 1

	index, found := binCode(a, target, 0, len(a)-1)
	fmt.Println(index, found)
}

func binCode(a []int, target, low, high int) (index int, found bool) {

	mid := (low + high) / 2

	if low > high {
		index = -1
		found = false
	} else {
		if target < a[mid] {
			index, found = binCode(a, target, low, mid-1)
		} else if target > a[mid] {
			index, found = binCode(a, target, mid+1, high)
		} else if target == a[mid] {
			index = mid
			found = true
		} else {
			index = -1
			found = false
		}
	}
	return
}
