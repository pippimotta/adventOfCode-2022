package main

import (
	_ "embed"
	"fmt"
	"strings"
)

const groupNumber = 3

//go:embed input.txt
var input string

func main() {
	itemsInAll := getPrimeItem(input)

	stringPrimaryItem := ""
	sum := 0

	for _, item := range itemsInAll {
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

func checkMap(s string) map[string]bool {
	check := map[string]bool{}
	for _, letter := range strings.Split(s, "") {
		check[letter] = true
	}
	return check
}

func getPrimeItem(input string) []string {
	var primeItem []string
	items := strings.Split(input, "\n")

	for i := 0; i < len(items); i += 3 {
		second, third := checkMap(items[i+1]), checkMap(items[i+2])
		for _, letter := range strings.Split(items[i], "") {
			if second[letter] && third[letter] {
				primeItem = append(primeItem, letter)
				break
			}
		}
	}
	return primeItem
}
