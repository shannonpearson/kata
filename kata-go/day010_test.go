package kata

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
)

var (
	oneByOne     = [][]int{{1}}
	twoByTwo     = [][]int{{1, 2}, {4, 3}}
	threeByThree = [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	fourByFour   = [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}
)

type createSpiralTestCase struct {
	n        int
	expected [][]int
}

var createSpiralTests = []createSpiralTestCase{
	{0, [][]int{}},
	{1, oneByOne},
	{2, twoByTwo},
	{3, threeByThree},
	{4, fourByFour},
}

// TestCreateSpiral tests kata CreateSpiral solution
func TestCreateSpiral(t *testing.T) {
	for _, c := range createSpiralTests {
		v := CreateSpiral(c.n)
		eq := true
		if v == nil || c.expected == nil {
			if len(v) != len(c.expected) {
				eq = false
			}
		} else if v != nil {
			for i, arr := range v {
				for j, num := range arr {
					if c.expected[i][j] != num {
						eq = false
					}
				}
			}
		}
		if !eq {
			t.Errorf("For %d expected %v got %v", colors.BrightYellow(c.n), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
