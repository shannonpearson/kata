package main

import "fmt"

// PosPeaks hold positions and values of peaks.
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// PickPeaks returns positions and values of local maxima.
func PickPeaks(array []int) PosPeaks {
	peaks := PosPeaks{Pos: []int{}, Peaks: []int{}}

	index := 0
	for i, num := range array {
		if i > 0 && i < len(array) {
			if num > array[i-1] {
				index = i
			} else if index != 0 && num < array[i-1] {
				peaks.Pos = append(peaks.Pos, index)
				peaks.Peaks = append(peaks.Peaks, array[index])
				index = 0
			}
		}
	}

	return peaks
}

func main() {
	input := []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}
	fmt.Println(PickPeaks(input))
}
