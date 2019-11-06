package main

import (
	"testing"
)

func TestBinCode(t *testing.T) {

	tests := []struct {
		input    []int
		target   int
		want     int
		wantedOK bool
	}{
		{
			[]int{1, 3, 4, 6, 8, 16, 17, 31, 33, 39},
			4,
			2,
			true,
		},
		{
			[]int{1, 3, 4, 6, 8, 16, 17, 31, 33, 39},
			7,
			5,
			false,
		},
		{
			[]int{1, 3, 4, 6, 8, 16, 17, 31, 33, 39},
			39,
			9,
			true,
		},
	}

	for _, test := range tests {
		got, ok := binCode(test.input, test.target, 0, len(test.input)-1)
		if got != test.want || ok != test.wantedOK {
			t.Errorf("expected (%d, %v), got (%d, %v)", test.want, test.wantedOK, got, ok)
		}
	}
}
