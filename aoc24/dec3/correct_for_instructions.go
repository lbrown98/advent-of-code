package dec3

import (
	"regexp"
	"strconv"
)

var (
	regex = `(mul\(([0-9]{1,3}),([0-9]{1,3})\))|(do\(\))|don't\(\)`
	numRegex = `([0-9]{1,3})`
	doRegex = `do\(\)`
	dontRegex = `don't\(\)`
)

func ParseCorruptedMemoryWithInstructions(input string) int {
	matches := matchMultipliesAndInstructions(input)
	return calculationWithInstructions(matches)
}

func matchMultipliesAndInstructions(input string) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(input, -1)
}

func calculationWithInstructions(matches []string) int {
	r := regexp.MustCompile(numRegex)
	do := regexp.MustCompile(doRegex)
	dont := regexp.MustCompile(dontRegex)

	instruction := true
	sum := 0

	for _, match := range matches {
		// check if it's a do/don't, update as necessary
		if do.FindString(match) != "" {
			instruction = true
		} else if dont.FindString(match) != ""{
			instruction = false
		} else if instruction {
			numbers := r.FindAllString(match, -1)
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			sum += num1 * num2
		}
	}
	return sum
}
