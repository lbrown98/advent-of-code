package dec1

import "sort"

func SimilarityScore(team_a []int, team_b []int) int {
	sort.Ints(team_a)
	sort.Ints(team_b)

	bListOccurances := map[int]int{}
	similarity := 0

	for _, b := range team_b {
		val, ok := bListOccurances[b]
		if !ok {
			bListOccurances[b] = 1
		} else {
			bListOccurances[b] = val + 1
		}
	}

	for _, a := range team_a {
		val, ok := bListOccurances[a]
		if ok {
			addToTotal := a * val
			similarity = similarity + addToTotal
		}
	}

	return similarity
}
