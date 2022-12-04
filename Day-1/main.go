package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	max, err := maxCal(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(max)
}
func maxCal(r io.Reader) (int, error) {
	max, elf := 0, 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			elf = 0
			continue
		}
		i, err := strconv.Atoi(text)
		if err != nil {
			return 0, err
		}
		elf += i
		if max < elf {
			max = elf
		}

	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return max, nil
}
