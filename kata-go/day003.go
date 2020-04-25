package kata

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// DigPow does a thing
func DigPow(n, p int) int {
	num := strconv.Itoa(n)
	digits := strings.Split(num, "")

	sum := 0
	for i, d := range digits {
		d, err := strconv.Atoi(d)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		sum += int(math.Pow(float64(d), float64(i+p)))
	}
	result := sum % n
	if result == 0 {
		return sum / n
	}
	return -1
}
