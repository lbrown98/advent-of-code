package dec2

import "math"

type Direction int

const (
	Ascending Direction = iota
	Descending
)

func SafeReports(reports [][]int, dampened bool) int {
	safeReports := 0

	for _, report := range(reports){
		if !dampened {
			safeReports += checkSafe(report)
		} else {
			removed := false
			safe := checkSafeDampened(report, removed)
			safeReports += safe
		}
	}
	
	return safeReports
}

func checkSafe(report []int) int {
	direction := Ascending

	if report[0] - report[1] > 0 {
		direction = Descending
	} else if report[0] - report[1] == 0 {
		return 0
	}

	for i := range report[:len(report)-1] {
		if direction == Ascending {
			if report[i] - report[i+1] >= 0 {
				return 0
			}
		} else {
			if report[i] - report[i+1] <= 0 {
				return 0
			}
		}
		if math.Abs(float64(report[i] - report[i+1])) > 3 {
			return 0
		}
	}

	return 1
}