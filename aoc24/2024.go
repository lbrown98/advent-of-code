package aoc24

import (
	"fmt"

	"github.com/lbrown98/advent-of-code/aoc24/config"
	"github.com/lbrown98/advent-of-code/aoc24/dec1"
	"github.com/lbrown98/advent-of-code/aoc24/dec2"
	"github.com/lbrown98/advent-of-code/aoc24/dec3"
	"github.com/lbrown98/advent-of-code/aoc24/dec4"
)

func day4() {
	totalXmas := dec4.WordSearch(config.WordSearch)
	fmt.Println("total: ", totalXmas)
}

func day3() {
	total := dec3.ParseCorruption(config.CorruptMemory)
	fmt.Println("total: ", total)

	total = dec3.ParseCorruptedMemoryWithInstructions(config.CorruptMemory)
	fmt.Println("total: ", total)
}

func day2() {
	safeReports := dec2.SafeReports(config.Levels, false)
	fmt.Println("safe reports: ", safeReports)

	dampenedSafeReports := dec2.SafeReports(config.Levels, true)
	fmt.Println("dampened safe reports: ", dampenedSafeReports)
}

func day1() {
	distance := dec1.DistanceSum(config.Team_a, config.Team_b)
	fmt.Println("distance: ", distance)

	score := dec1.SimilarityScore(config.Team_a, config.Team_b)
	fmt.Println("similarity score: ", score)
}
