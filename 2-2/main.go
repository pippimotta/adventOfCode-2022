package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	Win  = 6
	Lose = 0
	Draw = 3

	Scissor = 3
	Paper   = 2
	Rock    = 1
)

func main() {
	play, err := getPlay(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	sumScore := calScore(play)
	fmt.Printf("Your final score is %v\n", sumScore)
}

func getPlay(r io.Reader) ([][]string, error) {
	var playLine [][]string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		for _, play := range strings.Split(line, "\n") {
			playLine = append(playLine, strings.Split(play, " "))
		}

	}
	if err := scanner.Err(); err != nil {
		return [][]string{}, err
	}
	return playLine, nil
}

// X=Lose Y=Draw Z=win
func calScore(playLine [][]string) int {

	//set a hashMap for selecting your score corresponding to the result
	scoreMap := map[string]int{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}

	sumScore := 0
	for _, play := range playLine {
		//first add the score for result
		sumScore += scoreMap[play[1]]

		//then add the score for your choice of play
		switch play[0] {
		case "A": //op rock
			switch play[1] {
			case "X": //lose -> your choice is scissor
				sumScore += Scissor
			case "Y": //draw -> need rock
				sumScore += Rock
			case "Z": //win -> need paper
				sumScore += Paper
			}
		case "B": //op paper
			switch play[1] {
			case "X": //lose -> need rock
				sumScore += Rock
			case "Y": //draw -> need paper
				sumScore += Paper
			case "Z": //win -> need scissor
				sumScore += Scissor
			}
		case "C": //op scissor
			switch play[1] {
			case "X": //lose -> need paper
				sumScore += Paper
			case "Y": //draw -> need scissor
				sumScore += Scissor
			case "Z": //win -> need rock
				sumScore += Rock
			}
		}
	}
	return sumScore

}
