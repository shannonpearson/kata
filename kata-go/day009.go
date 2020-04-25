package kata

import "math"

func isPrime(num int) bool {
	sqrt := math.Sqrt(float64(num))
	i := 2
	for i <= int(sqrt) {
		if num%i == 0 {
			return false
		}
		i++
	}
	return true
}

func reverseInt(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}

// BackwardsPrime finds all backwards read primes between start and stop
func BackwardsPrime(start int, stop int) []int {
	backwardsPrimes := []int{}
	backwards := 0
	for i := start; i <= stop; i++ {
		backwards = reverseInt(i)
		if i != backwards && isPrime(i) && isPrime(backwards) {
			backwardsPrimes = append(backwardsPrimes, i)
		}
	}
	if len(backwardsPrimes) == 0 {
		return nil
	}
	return backwardsPrimes
}
