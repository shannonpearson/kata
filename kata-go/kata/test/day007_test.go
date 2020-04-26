package kata

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
	"github.com/shannonpearson/kata/kata-go/kata/kata"
)

type maximumSubarraySumTestCase struct {
	numbers  []int
	expected int
}

var maximumSubarraySumTests = []maximumSubarraySumTestCase{
	{[]int{}, 0},
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}, 0},
}

// TestMaximumSubarraySum tests kata MaximumSubarraySum solution
func TestMaximumSubarraySum(t *testing.T) {
	for _, c := range maximumSubarraySumTests {
		v := kata.MaximumSubarraySum(c.numbers)
		if v != c.expected {
			t.Errorf("For %v expected %d got %d", colors.BrightYellow(c.numbers), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
