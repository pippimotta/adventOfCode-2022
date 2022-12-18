package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	allWork := getWorkLines(input)
	overLapNum := 0
	for _, l := range allWork {
		if checkOverLap(l[:2], l[2:]) || checkOverLap(l[2:], l[:2]) {
			overLapNum++
		}
	}
	fmt.Printf("The number is %v\n", overLapNum)
}

func getWorkLines(input string) [][]int {
	var workLines [][]int
	for _, line := range strings.Split(input, "\n") {
		loads := strings.Split(line, ",")
		left := strings.Split(loads[0], "-")
		right := strings.Split(loads[1], "-")
		workLines = append(workLines, []int{stringToInt(left[0]), stringToInt(left[1]), stringToInt(right[0]), stringToInt(right[1])})
	}
	return workLines
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}

func checkOverLap(a, b []int) bool {

	return a[0] >= b[0] && a[1] <= b[1]
}
