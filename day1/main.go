package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func process() ([]int, []int) {
	data, err := os.ReadFile("input.txt")
	check(err)
	dataList := strings.Split(string(data), "\n")
	leftList, rightList := make([]int, len(dataList)), make([]int, len(dataList))
	for index, pair := range dataList {
		splitPair := strings.Split(pair, "   ")
		left, err := strconv.Atoi(splitPair[0])
		check(err)
		right, err := strconv.Atoi(splitPair[1])
		check(err)

		leftList[index] = int(left)
		rightList[index] = int(right)
	}
	return leftList, rightList

}

func part_one() int {
	result := 0
	leftList, rightList := process()

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i := 0; i < len(leftList); i++ {
		result += int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
	}

	return result

}

func part_two() int {
	result := 0
	leftList, rightList := process()
	value_count := make(map[int]int)

	for _, val := range rightList {
		value_count[val] += 1
	}

	for _, val := range leftList {
		result += val * value_count[val]
	}
	return result
}

func main() {
	fmt.Println("Part One:", part_one())
	fmt.Println("Part Two:", part_two())

}
