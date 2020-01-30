package kata_test

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
	"github.com/shannonpearson/kata-go/kata"
)

type gpsTestCase struct {
	s        int
	x        []float64
	expected int
}

var gpsTests = []gpsTestCase{
	{20, []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}, 41},
	{12, []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25}, 219},
}

// TestGps tests kata Gps solution
func TestGps(t *testing.T) {
	for _, c := range gpsTests {
		v := kata.Gps(c.s, c.x)
		if v != c.expected {
			t.Errorf("For %d, %v expected %v got %v", colors.BrightYellow(c.s), colors.BrightYellow(c.x), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
