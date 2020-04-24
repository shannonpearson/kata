package kata_test

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
	"github.com/shannonpearson/kata-go/kata"
)

type pickPeaksTestCase struct {
	array    []int
	expected kata.PosPeaks
}

var pickPeaksTests = []pickPeaksTestCase{
	{[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}, kata.PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}},
	{[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}, kata.PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}},
	{[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}, kata.PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}}},
	{[]int{2, 1, 3, 1, 2, 2, 2, 2, 1}, kata.PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}}},
	{[]int{2, 1, 3, 1, 2, 2, 2, 2}, kata.PosPeaks{Pos: []int{2}, Peaks: []int{3}}},
}

// TestPickPeaks tests kata PickPeaks solution
func TestPickPeaks(t *testing.T) {
	for _, c := range pickPeaksTests {
		v := kata.PickPeaks(c.array)
		eq := true
		for i, num := range v.Pos {
			if num != c.expected.Pos[i] || v.Peaks[i] != c.expected.Peaks[i] {
				eq = false
			}
		}

		if !eq {
			t.Errorf("For %v expected %v got %v", colors.BrightYellow(c.array), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
