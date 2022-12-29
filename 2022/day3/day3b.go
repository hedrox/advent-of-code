package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func convertCharToInt(char string) int {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(chars, char) + 1
}

func intersection(rucksacks []string) (map[string]bool, error) {
	if len(rucksacks) != 3 {
		return nil, errors.New("No of rucksacks != 3")
	}
	rucksackMap1 := make(map[rune]bool)
	rucksackMap2 := make(map[rune]bool)
	result := make(map[string]bool)

	for _, item := range rucksacks[0] {
		rucksackMap1[item] = true
	}
	for _, item := range rucksacks[1] {
		rucksackMap2[item] = true
	}

	for _, item := range rucksacks[2] {
		if _, ok := rucksackMap1[item]; ok {
			if _, ok := rucksackMap2[item]; ok {
				result[string(item)] = true
			}
		}
	}
	return result, nil
}

func main() {
	var rucksacks []string
	var prioritySum int

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		rucksacks = append(rucksacks, txt)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := 0; i < len(rucksacks); i += 3 {
		intersect, err := intersection(rucksacks[i : i+3])
		if err != nil {
			fmt.Println(err)
			continue
		}
		for key := range intersect {
			prioritySum += convertCharToInt(key)
		}
	}
	fmt.Println(prioritySum)
}
