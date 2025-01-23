package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("Day Three Solutions:")
	fmt.Println("Part One:", partOne())
	// partOne()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processInput() string {
	data, err := os.ReadFile("day3/input.txt")
	check(err)
	return string(data)
}

func partOne() int {
	corruptedData := processInput()
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	matched := re.FindAllString(corruptedData, -1)

	result := 0

	for _, match := range matched {
		numbers := match[4 : len(match)-1]
		firstNumber, err := strconv.Atoi(strings.Split(numbers, ",")[0])
		check(err)
		secondNumber, err := strconv.Atoi(strings.Split(numbers, ",")[1])
		check(err)
		result += firstNumber * secondNumber
	}
	return result
}
