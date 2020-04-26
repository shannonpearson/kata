package kata

// CreateSpiral creates an NxN spiral
func CreateSpiral(n int) [][]int {
	result := [][]int{}
	if n < 1 {
		return nil
	}
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}

	i := 1

	minRow := 0
	maxRow := n - 1
	minCol := 0
	maxCol := n - 1

	row := 0
	col := 0

	dir := "right"

	for i <= n*n {
		if dir == "right" {
			for col <= maxCol {
				result[row][col] = i
				i++
				col++
			}
			minRow++
			dir = "down"
			col--
			row++
		} else if dir == "left" {
			for col >= minCol {
				result[row][col] = i
				i++
				col--
			}
			maxRow--
			dir = "up"
			col++
			row--
		} else if dir == "down" {
			for row <= maxRow {
				result[row][col] = i
				i++
				row++
			}
			maxCol--
			dir = "left"
			row--
			col--
		} else {
			for row >= minRow {
				result[row][col] = i
				i++
				row--
			}
			minCol++
			dir = "right"
			row++
			col++
		}
	}
	return result
}
