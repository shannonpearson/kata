package kata

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
)

type backwardsPrimeTestCase struct {
	start    int
	stop     int
	expected []int
}

var backwardsPrimeTests = []backwardsPrimeTestCase{
	{1, 100, []int{13, 17, 31, 37, 71, 73, 79, 97}},
	{1, 31, []int{13, 17, 31}},
	{501, 599, nil},
}

// TestBackwardsPrime tests kata BackwardsPrime solution
func TestBackwardsPrime(t *testing.T) {
	for _, c := range backwardsPrimeTests {
		v := BackwardsPrime(c.start, c.stop)
		eq := true
		if (c.expected == nil && v != nil) || (v == nil && c.expected != nil) {
			eq = false
		} else {
			for i, num := range v {
				if num != c.expected[i] {
					eq = false
				}
			}
		}
		if !eq {
			t.Errorf("For %d, %d expected %v got %v", colors.BrightYellow(c.start), colors.BrightYellow(c.stop), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
