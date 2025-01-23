package day2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("Day Two Solutions:")
	fmt.Println("Part One:", partOne())
	fmt.Println("Part Two:", partTwo())

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type DeltaState int

const (
	Undefined  DeltaState = 0
	Increasing DeltaState = 1
	Decreasing DeltaState = -1
)

func process() [][]int {
	data, err := os.ReadFile("/Users/maxwellgonick/dev/advent_of_code/day2/input.txt")
	check(err)
	dataList := strings.Split(string(data), "\n")
	reportList := make([][]int, len(dataList))
	for reportIndex, report := range dataList {
		reportString := strings.Fields(report)
		reportList[reportIndex] = make([]int, len(reportString))
		for valIndex, val := range reportString {
			valInt, err := strconv.Atoi(val)
			check(err)
			reportList[reportIndex][valIndex] = valInt
		}
	}
	return reportList

}

func processReport(report []int) bool {
	var state = Undefined
	delta := report[0] - report[1]
	switch {
	case delta > 0:
		state = Decreasing
	case delta < 0:
		state = Increasing
	default:
		return false
	}
	for i := 0; i < len(report)-1; i++ {
		delta := report[i] - report[i+1]
		localState := DeltaState(Undefined)
		switch {
		case delta > 0:
			localState = Decreasing
		case delta < 0:
			localState = Increasing
		}

		if localState != state {
			return false
		}

		if math.Abs(float64(delta)) > 3 {
			return false
		}
	}
	return true
}

func partOne() int {
	reportList := process()
	result := len(reportList)
	for _, report := range reportList {
		reportResult := processReport(report)
		if !reportResult {
			result -= 1
		}
	}
	return result
}

func partTwo() int {
	goodReports := 0
	reportList := process()
	for _, report := range reportList {

		if processReport(report) {
			goodReports++
			continue
		}
		for index := range report {
			newReport := make([]int, len(report)-1)
			copy(newReport[:index], report[:index])
			copy(newReport[index:], report[index+1:])

			if processReport(newReport) {
				goodReports++
				break
			}

		}
	}
	return goodReports
}
