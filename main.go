package main

import (
	"fmt"

	"github.com/lbrown98/advent-of-code/2024/dec1"
)

func main() {
	day1()
}

func day1() {
	distance := dec1.DistanceSum(team_a, team_b)
	fmt.Println("distance: ", distance)

	score := dec1.SimilarityScore(team_a, team_b)
	fmt.Println("similarity score: ", score)
}