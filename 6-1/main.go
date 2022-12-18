package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type set map[string]struct{}

func (s set) Add(value string) {
	s[value] = struct{}{}
}

func (s set) Contains(value string) bool {
	_, ok := s[value]
	return ok
}

func (s set) Remove(value string) {
	delete(s, value)
}

func (s set) Len() int {
	return len(s)
}

func main() {
	var existMap = set{}
	message := strings.Split(input, "")
	l, counter := 0, 0

	for r := 0; r < len(message); r++ {
		counter++
		for existMap.Contains(message[r]) {
			existMap.Remove(message[l])
			l++
		}
		existMap.Add(message[r])
		if existMap.Len() == 4 { //14 for part2
			fmt.Printf("%d characters need to be processed\n", counter)
			break
		}
	}

}
