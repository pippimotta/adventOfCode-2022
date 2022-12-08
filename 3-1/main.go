package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	primaryItem := getPrimaryItem(getSackLine(input))

	stringPrimaryItem := ""
	sum := 0

	for _, item := range primaryItem {
		stringPrimaryItem += item
	}
	for _, i := range stringPrimaryItem {
		value := int(i)
		sum += value
		if value >= 65 && value <= 90 { //uppercase
			sum -= 38
		} else {
			sum -= 96
		}

	}
	fmt.Printf("The sum of priorities is %d\n", sum)

}

func getSackLine(s string) [][]string {
	var sackLine [][]string
	for _, l := range strings.Split(s, "\n") {
		sackLine = append(sackLine, getIndividualCompart(l))
	}
	return sackLine
}

func getIndividualCompart(s string) []string {
	var splitedCompart []string
	splitedCompart = append(splitedCompart, s[:len(s)/2], s[len(s)/2:])
	return splitedCompart
}

func wordCount(s string) map[string]int {
	countMap := make(map[string]int)
	for _, str := range strings.Split(s, "") {
		countMap[str]++
	}
	return countMap
}

func getCommonLetter(a, b string) string {
	countA := wordCount(a)
	for _, str := range strings.Split(b, "") {
		if _, ok := countA[str]; ok {
			return str
		}
	}
	return ""
}

func getPrimaryItem(s [][]string) []string {
	var primaryItem []string
	for _, itemLine := range s {
		primaryItem = append(primaryItem, getCommonLetter(itemLine[0], itemLine[1]))
	}
	return primaryItem
}
