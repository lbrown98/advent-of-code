package dec2

import "math"

type Direction int

const (
	Unsafe Direction = iota
	Ascending
	Descending
)

func SafeReports(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		safe := false

		direction := checkDirection(report[0], report[1])
		if direction != Unsafe {
			safe = assessReports(report, direction)
		}

		if safe {
			safeReports = safeReports + 1
		}
	}

	return safeReports
}

func assessReports(report []int, direction Direction) bool {
	prev := -1

	for _, level := range report {
		if prev != -1 {
			change := float64(level-prev)
			if math.Abs(change) > 3 || math.Abs(change) < 1 {
				return false
			}
			if change > 0 && direction == Descending {
				return false
			}
			if change < 0 && direction == Ascending {
				return false
			}
		}
		prev = level

	}

	return true
}

func checkDirection(a int, b int) Direction {
	var direction Direction
	change := b - a

	if change > 0 {
		direction = Ascending
	} else if change < 0 {
		direction = Descending
	} else {
		direction = Unsafe
	}

	return direction
}
