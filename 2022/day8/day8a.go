package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	cornerTrees = 4
)

func treesAtMargin(treeGrid [][]int) int {
	// We already include all the corner trees in the first multiplication
	// we subtract them from the second multiplication
	return (len(treeGrid) * 2) + ((len(treeGrid[0]) * 2) - cornerTrees)
}

func getVisibleTrees(treeGrid [][]int, row, col int) int {
	i := row
	iVisible := true
	for i > 0 {
		i--
		if treeGrid[i][col] >= treeGrid[row][col] {
			iVisible = false
		}
	}
	if iVisible {
		return 1
	}
	iVisible = true
	i = row
	for i < len(treeGrid)-1 {
		i++
		if treeGrid[i][col] >= treeGrid[row][col] {
			iVisible = false
		}
	}
	if iVisible {
		return 1
	}

	j := col
	jVisible := true
	for j > 0 {
		j--
		if treeGrid[row][j] >= treeGrid[row][col] {
			jVisible = false
		}
	}
	if jVisible {
		return 1
	}
	jVisible = true
	j = col
	for j < len(treeGrid[0])-1 {
		j++
		if treeGrid[row][j] >= treeGrid[row][col] {
			jVisible = false
		}
	}
	if jVisible {
		return 1
	}
	return 0
}

func main() {
	var visibleTrees int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	treeGrid := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		row := make([]int, 0)
		for _, char := range txt {
			if ch, err := strconv.Atoi(string(char)); err == nil {
				row = append(row, ch)
			}
		}
		treeGrid = append(treeGrid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	visibleTrees += treesAtMargin(treeGrid)
	for i := 1; i < len(treeGrid)-1; i++ {
		for j := 1; j < len(treeGrid[i])-1; j++ {
			visibleTrees += getVisibleTrees(treeGrid, i, j)
		}
	}
	fmt.Println(visibleTrees)
}
