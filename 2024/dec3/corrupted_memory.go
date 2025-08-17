package dec3

import (
	"regexp"
	"strconv"
)

func ParseCorruption(input string) int {
	matches := matchMultiplies(input)
	return calculation(matches)
}

func matchMultiplies(input string) []string {
	regex := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`
	r := regexp.MustCompile(regex)
	return r.FindAllString(input, -1)
}

func calculation(matches []string) int {
	regex := `([0-9]{1,3})`
	r := regexp.MustCompile(regex)
	sum := 0

	for _, mult := range matches {
		numbers := r.FindAllString(mult, -1)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		sum += num1 * num2
	}
	return sum
}
