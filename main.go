package main

import (
	"fmt"

	"github.com/lbrown98/advent-of-code/2024/dec1"
	"github.com/lbrown98/advent-of-code/2024/dec2"
)

func main() {
	day2()
}

func day2() {
	safeReports := dec2.SafeReports(levels)
	fmt.Println("safe reports: ", safeReports)

	// dampenedSafeReports := dec2.Dampener(levels)
	// fmt.Println("dampened safe reports: ", dampenedSafeReports)
}

func day1() {
	distance := dec1.DistanceSum(team_a, team_b)
	fmt.Println("distance: ", distance)

	score := dec1.SimilarityScore(team_a, team_b)
	fmt.Println("similarity score: ", score)
}
