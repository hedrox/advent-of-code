package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCycle(input []string, idx int) int {
	xRegister := 1
	for i, command := range input {
		if idx == i {
			break
		}
		if strings.HasPrefix(command, "addx") {
			cmd := strings.Split(command, " ")
			value, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}
			xRegister += value
		}
	}
	return xRegister
}

func main() {
	var input []string
	var sumOfSignals int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "noop" {
			input = append(input, txt)
		} else {
			// We add a noop before an addx command so that it will take 2 cycles
			input = append(input, "noop")
			input = append(input, txt)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	cycles := []int{20, 60, 100, 140, 180, 220}
	for _, cycle := range cycles {
		sumOfSignals += cycle * getCycle(input, cycle-1)
	}
	fmt.Println(sumOfSignals)
}
