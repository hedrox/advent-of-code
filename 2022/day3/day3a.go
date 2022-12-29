package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertCharToInt(char string) int {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(chars, char) + 1
}

func intersection(a, b string) map[string]bool {
	aMap := make(map[rune]bool)
	result := make(map[string]bool)

	for _, item := range a {
		aMap[item] = true
	}

	for _, item := range b {
		if _, ok := aMap[item]; ok {
			result[string(item)] = true
		}
	}
	return result
}

func main() {
	var prioritySum int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		mid := len(txt) / 2
		leftCompartment := txt[:mid]
		rightCompartment := txt[mid:]
		intersect := intersection(leftCompartment, rightCompartment)
		for char := range intersect {
			prioritySum += convertCharToInt(char)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(prioritySum)
}
