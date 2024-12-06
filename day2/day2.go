package day2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("Day One Solutions:")
	fmt.Println("Part One:", part_one())
	// fmt.Println("Part Two:", part_two())

	part_one()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func process() [][]int {
	data, err := os.ReadFile("day2/input.txt")
	check(err)
	dataList := strings.Split(string(data), "\n")
	reportList := make([][]int, len(dataList))
	for reportIndex, report := range dataList {
		reportString := strings.Fields(report)
		reportList[reportIndex] = make([]int, len(reportString))
		for valIndex, val := range reportString {
			val_int, err := strconv.Atoi(val)
			check(err)
			reportList[reportIndex][valIndex] = val_int
		}
	}
	return reportList

}

func part_one() int {
	type DeltaState int

	const (
		Undefined  DeltaState = 0
		Increasing DeltaState = 1
		Decreasing DeltaState = -1
	)

	reportList := process()
	result := len(reportList)
	for _, report := range reportList {
		var state DeltaState = Undefined
		for i := 0; i < len(report)-1; i++ {
			delta := report[i+1] - report[i]
			// Set State
			if delta > 0 {
				if state == Undefined {
					state = Increasing
				}
				if state == Decreasing {
					result -= 1
					break
				}
			} else if delta < 0 {
				if state == Undefined {
					state = Decreasing
				}
				if state == Increasing {
					result -= 1
					break
				}
			} else {
				//Two Same Levels
				result -= 1
				break
			}

			if math.Abs(float64(delta)) > 3 {
				result -= 1
				break
			}
		}
	}
	return result
}
