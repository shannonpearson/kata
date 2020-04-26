package kata

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
)

type bouncingBallTestCase struct {
	height   float64
	bounce   float64
	window   float64
	expected int
}

var bouncingBallTests = []bouncingBallTestCase{
	{3, 0.66, 1.5, 3},
	{40, 0.4, 10, 3},
	{10, 0.6, 10, -1},
	{40, 1, 10, -1},
	{5, -1, 1.5, -1},
}

// TestBouncingBall tests kata BouncingBall solution
func TestBouncingBall(t *testing.T) {
	for _, c := range bouncingBallTests {
		v := BouncingBall(c.height, c.bounce, c.window)
		if v != c.expected {
			t.Errorf("For %f, %f, %f expected %d got %d", colors.BrightYellow(c.height), colors.BrightYellow(c.bounce), colors.BrightYellow(c.window), colors.BrightGreen(c.expected), colors.BrightRed(v))
		}
	}
}
