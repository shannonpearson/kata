package kata

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
)

type countBitsTestCase struct {
	num      uint
	expected int
}

var countBitsTests = []countBitsTestCase{
	{0, 0},
	{4, 1},
	{7, 3},
	{9, 2},
	{10, 2},
}

// TestCountBits tests kata CountBits solution
func TestCountBits(t *testing.T) {
	for _, c := range countBitsTests {
		v := CountBits(c.num)
		if v != c.expected {
			t.Errorf("For %d expected %d got %d", colors.BrightYellow(c.num), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
