package kata

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
)

type twoSumTestCase struct {
	numbers  []int
	target   int
	expected [2]int
}

var twoSumTests = []twoSumTestCase{
	{[]int{1, 2, 3}, 4, [2]int{0, 2}},
	{[]int{1234, 5678, 9012}, 14690, [2]int{1, 2}},
	{[]int{2, 2, 3}, 4, [2]int{0, 1}},
}

// TestTwoSum tests kata TwoSum solution
func TestTwoSum(t *testing.T) {
	for _, c := range twoSumTests {
		v := TwoSum(c.numbers, c.target)
		if v != c.expected {
			t.Errorf("For %v, %d expected %v got %v", colors.BrightYellow(c.numbers), colors.BrightYellow(c.target), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
