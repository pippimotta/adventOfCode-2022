package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var stacks = make(map[int][]rune)
	var cranes = [][]rune{
		{'R', 'G', 'H', 'Q', 'S', 'B', 'T', 'N'},
		{'H', 'S', 'F', 'D', 'P', 'Z', 'J'},
		{'Z', 'H', 'V'},
		{'M', 'Z', 'J', 'F', 'G', 'H'},
		{'T', 'Z', 'C', 'D', 'L', 'M', 'S', 'R'},
		{'M', 'T', 'W', 'V', 'H', 'Z', 'J'},
		{'T', 'F', 'P', 'L', 'Z'},
		{'Q', 'V', 'W', 'S'},
		{'W', 'H', 'L', 'M', 'T', 'D', 'N', 'C'},
	}

	for i := 1; i < 10; i++ {
		stacks[i] = cranes[i-1]
	}
	moves := strings.Split(input, "\n")
	for _, move := range moves {
		var step int
		var from int
		var to int
		fmt.Sscanf(move, "move %d from %d to %d", &step, &from, &to)
		source := stacks[from]
		destination := stacks[to]

		target := len(source) - step
		for _, c := range source[target:] {
			destination = append(destination, c)
		}
		source = source[:target]
		stacks[from] = source
		stacks[to] = destination
	}
	finalCrane := ""
	for j := 1; j < 10; j++ {
		finalCrane += string(stacks[j][len(stacks[j])-1])
	}
	fmt.Println("Final cranes at the top are", finalCrane)

}
