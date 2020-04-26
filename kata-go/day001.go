package kata

// TwoSum takes an array of integers and a target numbers and returns the
// indices of the integers whose sum is equal to the target number.
func TwoSum(numbers []int, target int) [2]int {
	indices := [2]int{}
outer:
	for i, num := range numbers {
		for j, n := range numbers[i+1:] {
			if num+n == target {
				indices = [2]int{i, i + j + 1}
				break outer
			}
		}
	}
	return indices
}
