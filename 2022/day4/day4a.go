package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var noContained int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		text := strings.Split(txt, ",")
		var lhs0, lhs1, rhs0, rhs1 int
		fmt.Sscanf(text[0], "%d-%d", &lhs0, &lhs1)
		fmt.Sscanf(text[1], "%d-%d", &rhs0, &rhs1)

		if lhs0 >= rhs0 && lhs1 <= rhs1 || rhs0 >= lhs0 && rhs1 <= lhs1 {
			noContained++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(noContained)
}
