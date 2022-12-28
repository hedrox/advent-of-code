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
	responseChoice := map[string]string{"X": "rock", "Y": "paper", "Z": "scissors"}

	if opponentChoice[opponent] == responseChoice[response] {
		return shapeSelected[responseChoice[response]] + roundOutcomeScore["draw"]
	}
	if (opponentChoice[opponent] == "rock" && responseChoice[response] == "paper") ||
		(opponentChoice[opponent] == "paper" && responseChoice[response] == "scissors") ||
		(opponentChoice[opponent] == "scissors" && responseChoice[response] == "rock") {
		return shapeSelected[responseChoice[response]] + roundOutcomeScore["won"]
	}
	return shapeSelected[responseChoice[response]] + roundOutcomeScore["loss"]
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
