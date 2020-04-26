package kata

// BouncingBall calculates the number of times a ball bounces in front of a window.
func BouncingBall(height, bounce, window float64) int {
	if height <= 0 || bounce <= 0 || bounce >= 1 || window >= height {
		return -1
	}

	count := -1

	for ; height > window; height *= bounce {
		count += 2
	}

	return count
}
