package dec1

import (
	"math"
	"sort"
)

func DistanceSum(team_a []int, team_b []int) int {
	sort.Ints(team_a)
	sort.Ints(team_b)

	sum := 0.0

	for i := range team_a {
		add := math.Abs(float64(team_a[i] - team_b[i]))
		sum = sum + add
	}

	return int(sum)
}
