package main

import "fmt"

// MaximumSubarraySum returns the maximum sum of any set of contiguous values in a slice
func MaximumSubarraySum(numbers []int) int {
	max := 0
	for i, num := range numbers {
		sum := num
		if num > max {
			max = num
		}
		for _, n := range numbers[i+1:] {
			sum += n
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaximumSubarraySum(nums))
}
