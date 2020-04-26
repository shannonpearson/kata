package kata

func calcHourlySpeed(d float64, s int) float64 {
	return (3600.0 * d) / float64(s)
}

// Gps calculates the floor of average speed per hour
func Gps(s int, x []float64) int {
	if len(x) <= 1 {
		return 0
	}

	highest := 0.0
	h := 0.0
	for i, t := range x {
		if i != 0 {
			h = calcHourlySpeed(t-x[i-1], s)
			if h > highest {
				highest = h
			}
		}
	}

	return int(highest)
}
