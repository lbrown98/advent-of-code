package dec2

import (
	"math"
)

func Dampener(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		safe := false

		direction := checkDirection(report[0], report[1])
		if direction != Unsafe {
			safe = assessReportsWithDampener(report, direction)
		}

		if safe {
			safeReports = safeReports + 1
		}
	}

	return safeReports
}

func assessReportsWithDampener(report []int, direction Direction) bool {
	prev := -1
	dampenerUsed := false

	for i, level := range report {
		if prev == -1 {
			prev = level
		} else {
			change := float64(level-prev)
			if (math.Abs(change) > 3 || math.Abs(change) < 1) || (change > 0 && direction == Descending) ||
				(change < 0 && direction == Ascending) {
				if dampenerUsed {
					return false
				}
				dampenerUsed = true
				if i == 1 {
					direction = checkDirection(prev, report[2])
				}
			} else {
				prev = level
			}
		}

	}

	return true
}
