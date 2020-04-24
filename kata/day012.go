package kata

import (
	"fmt"
	"math"
	"sort"
)

// there is probably a faster way to find all the factors...
func getPrimeFactors(i int) []int {
	c := i

	factors := []int{}
	for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
		if i%j == 0 {
			factors = append(factors, getPrimeFactors(j)...)
			factors = append(factors, getPrimeFactors(i/j)...)
			i = i / j
		}
	}

	if i == c {
		factors = append(factors, i)
	}
	return factors
}

func SumOfDivided(lst []int) string {
	result := ""
	primes := []int{}
	dict := make(map[int]map[int]bool)

	for _, n := range lst {
		factors := getPrimeFactors(n)
		for _, f := range factors {
			if dict[f] == nil {
				dict[f] = make(map[int]bool)
				primes = append(primes, f)
			}
			dict[f][n] = true
		}

	}

	sort.Ints(primes)
	for _, value := range primes {
		sum := 0
		for k := range dict[value] {
			sum += k
		}
		result += fmt.Sprintf("(%d %d)", value, sum)
	}

	return result
}
