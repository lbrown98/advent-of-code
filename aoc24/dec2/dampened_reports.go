package dec2

import "math"

func checkSafeDampened(report []int, removed bool) int {
	direction := Ascending

	if report[0]-report[1] > 0 {
		direction = Descending
	} else if report[0]-report[1] == 0 {
		if !removed {
			if checkSafeDampened(removeIndex(report, 0), true)+
				checkSafeDampened(removeIndex(report, 1), true) >= 1 {
				return 1
			}
		}
		return 0
	}

	for i := range report[:len(report)-1] {
		if direction == Ascending {
			if report[i]-report[i+1] >= 0 {
				if !removed {
					if checkIfRemovingIndexZeroWorks(report) == 1 {
						return 1
					}
					if checkSafeDampened(removeIndex(report, i), true)+
						checkSafeDampened(removeIndex(report, i+1), true) >= 1 {
						return 1
					}
				}
				return 0
			}
		} else {
			if report[i]-report[i+1] <= 0 {
				if !removed {
					if checkIfRemovingIndexZeroWorks(report) == 1 {
						return 1
					}
					if checkSafeDampened(removeIndex(report, i), true)+
						checkSafeDampened(removeIndex(report, i+1), true) >= 1 {
						return 1
					}
				}
				return 0
			}
		}
		if math.Abs(float64(report[i]-report[i+1])) > 3 {
			if !removed {
				if checkIfRemovingIndexZeroWorks(report) == 1 {
					return 1
				}
				if checkSafeDampened(removeIndex(report, i), true)+
					checkSafeDampened(removeIndex(report, i+1), true) >= 1 {
					return 1
				}
			}
			return 0
		}
	}

	return 1
}

func removeIndex(report []int, indexToRemove int) []int {
	reportModified := make([]int, len(report))
	copy(reportModified, report)

	if indexToRemove == len(report)-1 {
		reportModified = reportModified[:len(report)-1]
	} else if indexToRemove >= 0 && indexToRemove < len(report)-1 {
		reportModified = append(reportModified[:indexToRemove], reportModified[indexToRemove+1:]...)
	}
	return reportModified
}

func checkIfRemovingIndexZeroWorks(report []int) int {
	return checkSafe(removeIndex(report, 0))
}