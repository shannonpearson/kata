package kata

import (
	"strconv"
)

// CountBits returns the number of 1 digits in the binary representation of an unsigned integer
func CountBits(num uint) int {
	n := uint64(num)
	bin := strconv.FormatUint(n, 2)
	count := 0
	for _, char := range bin {
		if char == '1' {
			count++
		}
	}
	return count
}

// func CountBits(n uint) int {
// 	return bits.OnesCount(n) // from math/bits packate
// 	return strings.Count(strconv.FormatInt(int64(n), 2), "1")
// }
