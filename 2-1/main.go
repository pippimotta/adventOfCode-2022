package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
)

var input string

func main() {
	sum, err := sumScore(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Your final score is %v\n", sum)
}

func sumScore(r io.Reader) (int, error) {
	sum := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		sum += judge(text)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}

// not sure who is the computer(´Д` )
func judge(s string) int {
	switch s {
	case "A X":
		return 4
	case "A Y":
		return 8
	case "A Z":
		return 3
	case "B X":
		return 1
	case "B Y":
		return 5
	case "B Z":
		return 9
	case "C X":
		return 7
	case "C Y":
		return 2
	case "C Z":
		return 6
	}
	return 0
}
