package main

import "fmt"

// FindOdd finds the integer in a slice of ints that appears an odd number of times
func FindOdd(seq []int) int {
	// your code here
	ints := map[int]bool{}
	for _, n := range seq {
		if val, ok := ints[n]; ok {
			ints[n] = !val
		} else {
			ints[n] = true
		}
	}
	for k, v := range ints {
		if v {
			return k
		}
	}
	return 0
}

func main() {
	input := []int{
		20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5,
	}
	fmt.Println(FindOdd(input))
}
