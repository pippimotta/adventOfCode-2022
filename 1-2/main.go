package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	maxThree, err := maxThreeCal(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("The top three Calorie in total is %v\n", maxThree)
}

func maxThreeCal(r io.Reader) (int, error) {
	elf := 0
	// (´Д` )
	var calPool []int
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			calPool = append(calPool, elf)

			elf = 0
			continue
		}

		i, err := strconv.Atoi(text)
		if err != nil {
			return 0, err
		}
		elf += i
	}
	calPool = append(calPool, elf)

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	sort.Slice(calPool, func(i, j int) bool {
		return calPool[i] > calPool[j]
	})
	var maxThree int
	for _, n := range calPool[:3] {
		maxThree += n
	}
	return maxThree, nil
}
