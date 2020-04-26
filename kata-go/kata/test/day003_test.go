package kata

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
	"github.com/shannonpearson/kata/kata-go/kata/kata"
)

type digPowTestCase struct {
	n        int
	p        int
	expected int
}

var digPowTests = []digPowTestCase{
	{89, 1, 1},
	{92, 1, -1},
	{695, 2, 2},
	{46288, 3, 51},
}

// TestDigPow tests kata DigPow solution
func TestDigPow(t *testing.T) {
	for _, c := range digPowTests {
		v := kata.DigPow(c.n, c.p)
		if v != c.expected {
			t.Errorf("For %v, %d expected %v got %v", colors.BrightYellow(c.n), colors.BrightYellow(c.p), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
