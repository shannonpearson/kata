package kata_test

import (
	"testing"

	colors "github.com/logrusorgru/aurora"
	"github.com/shannonpearson/kata-go/kata"
)

var passStrings = []string{
	"abcde",
	"abc123",
	"12345",
	"0",
	"ABcDe",
	"ABC",
}

var failStrings = []string{
	"",
	"abcd ",
	"ab+cd",
	"+",
	"ab_cd",
	"    ",
	"abc)",
	"$ab",
	"h%i",
	"3%4",
	".xyz",
	"ABc3@",
}

// TestAlphanumericTrue tests true alphanumeric
func TestAlphanumericTrue(t *testing.T) {
	for _, s := range passStrings {
		v := kata.Alphanumeric(s)
		if v != true {
			t.Errorf("For %v expected %v got %v", colors.BrightYellow(s), colors.BrightGreen("true"), colors.BrightRed("false"))
		}
	}
}

// TestAlphanumericFalse tests false alphanumeric
func TestAlphanumericFalse(t *testing.T) {
	for _, s := range failStrings {
		v := kata.Alphanumeric(s)
		if v == true {
			t.Errorf("For %v expected %v got %v", colors.BrightYellow(s), colors.BrightGreen("false"), colors.BrightRed("true"))
		}
	}
}
