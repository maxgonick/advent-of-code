package day3

import (
	"fmt"
	"os"
)

func Solve() {
	fmt.Println("Day Three Solutions:")
	//fmt.Println("Part One:", partOne())
	partOne()
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

func partOne() {
	corruptedData := processInput()

}
