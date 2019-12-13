package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	tests := []struct {
		input []int
		want  []int
	}{
		{
			[]int{23, 11, 43, 52, 22, 67, 56, 34, 18},
			[]int{11, 18, 22, 23, 34, 43, 52, 56, 67},
		},

		{
			[]int{},
			[]int{},
		},

		{
			[]int{5, 7, 4, 6, 3, 9, 2, 8, 1, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, test := range tests {
		BubbleSort(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("slice %v should be sorted as %v", test.input, test.want)
		}
	}
}
