package kata_test

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
	"github.com/shannonpearson/kata-go/kata"
)

var inputs = [...][]int{
	[]int{12, 15},
	[]int{15, 21, 24, 30, 45},
	[]int{15, -21, 24, 30, 45},
	[]int{15, 9, 37, 45},
	[]int{15, 9, 37, 9},
}

var solutions = []string{
	"(2 12)(3 27)(5 15)",
	"(2 54)(3 135)(5 90)(7 21)",
	"()",
	"(3 69)(5 60)(37 37)",
	"(3 78)(5 60)(37 37)",
}

// TestSumOfDivided tests SumOfDivided
func TestSumOfDivided(t *testing.T) {
	for i, a := range inputs {
		v := kata.SumOfDivided(a)
		if v != solutions[i] {
			t.Errorf("For %v expected %v got %v", colors.BrightYellow(inputs[i]), colors.BrightGreen(solutions[i]), colors.BrightRed(v))
		}
	}
}
