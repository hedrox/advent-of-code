package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcScore(opponent string, response string) int {
	shapeSelected := map[string]int{"rock": 1, "paper": 2, "scissors": 3}
	roundOutcomeScore := map[string]int{"loss": 0, "draw": 3, "won": 6}
	opponentChoice := map[string]string{"A": "rock", "B": "paper", "C": "scissors"}
	responseChoice := map[string]string{"X": "loss", "Y": "draw", "Z": "won"}

	if responseChoice[response] == "draw" {
		return shapeSelected[opponentChoice[opponent]] + roundOutcomeScore["draw"]
	}
	if responseChoice[response] == "won" {
		if opponentChoice[opponent] == "rock" {
			return shapeSelected["paper"] + roundOutcomeScore["won"]
		}
		if opponentChoice[opponent] == "paper" {
			return shapeSelected["scissors"] + roundOutcomeScore["won"]
		}
		if opponentChoice[opponent] == "scissors" {
			return shapeSelected["rock"] + roundOutcomeScore["won"]
		}
	} else {
		if opponentChoice[opponent] == "paper" {
			return shapeSelected["rock"] + roundOutcomeScore["loss"]
		}
		if opponentChoice[opponent] == "scissors" {
			return shapeSelected["paper"] + roundOutcomeScore["loss"]
		}
		if opponentChoice[opponent] == "rock" {
			return shapeSelected["scissors"] + roundOutcomeScore["loss"]
		}
	}
	return 0
}

func main() {
	var totalScore int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		text := strings.Split(txt, " ")
		totalScore += calcScore(text[0], text[1])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(totalScore)
}
