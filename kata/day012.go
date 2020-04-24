package kata

import (
	"fmt"
	"math"
	"sort"
)

// there is probably a faster way to find all the factors...
func getPrimeFactors(i int) []int {
	i = int(math.Abs(float64(i)))
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

	uniqDict := make(map[int]bool)
	for _, n := range factors {
		uniqDict[n] = true
	}

	uniq := []int{}
	for k := range uniqDict {
		uniq = append(uniq, k)
	}

	return uniq
}

func SumOfDivided(lst []int) string {
	result := ""
	primes := []int{}
	dict := make(map[int][]int)

	for _, n := range lst {
		factors := getPrimeFactors(n)
		for _, f := range factors {
			if dict[f] == nil {
				dict[f] = []int{}
				primes = append(primes, f)
			}
			dict[f] = append(dict[f], n)
		}

	}
	fmt.Println(dict)

	sort.Ints(primes)
	for _, value := range primes {
		sum := 0
		for _, m := range dict[value] {
			sum += m
		}
		result += fmt.Sprintf("(%d %d)", value, sum)
	}

	return result
}
