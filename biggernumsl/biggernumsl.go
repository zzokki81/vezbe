/*This package append all numbers bigger then n from
slice a to slice b */

package biggernumsl

func LargerNums(n int, a []int) []int {
	Temp := []int{}
	for _, e := range a {
		if e > n {

			Temp = append(Temp, e)
		}
	}
	return Temp
}
