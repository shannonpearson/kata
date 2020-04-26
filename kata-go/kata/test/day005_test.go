package kata

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
	"github.com/shannonpearson/kata/kata-go/kata/kata"
)

type findOddTestCase struct {
	seq      []int
	expected int
}

var findOddTests = []findOddTestCase{
	{[]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}, 5},
	{[]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}, -1},
	{[]int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}, 5},
	{[]int{10}, 10},
	{[]int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}, 10},
	{[]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}, 1},
}

// TestFindOdd tests kata FindOdd solution
func TestFindOdd(t *testing.T) {
	for _, c := range findOddTests {
		v := kata.FindOdd(c.seq)
		if v != c.expected {
			t.Errorf("For %v expected %d got %d", colors.BrightYellow(c.seq), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
