package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func treesAtMargin(treeGrid [][]int) int {
	// We already include all the corner trees in the first multiplication
	// we subtract them from the second multiplication
	return (len(treeGrid) * 2) + ((len(treeGrid[0]) * 2) - 4)
}

func getVisibleTrees(treeGrid [][]int, i, j int) int {
	x := i
	iVisible := true
	for x > 0 {
		x--
		if treeGrid[x][j] >= treeGrid[i][j] {
			iVisible = false
		}
	}
	if iVisible {
		return 1
	}
	iVisible = true
	x = i
	for x < len(treeGrid)-1 {
		x++
		if treeGrid[x][j] >= treeGrid[i][j] {
			iVisible = false
		}
	}
	if iVisible {
		return 1
	}

	y := j
	jVisible := true
	for y > 0 {
		y--
		if treeGrid[i][y] >= treeGrid[i][j] {
			jVisible = false
		}
	}
	if jVisible {
		return 1
	}
	jVisible = true
	y = j
	for y < len(treeGrid[0])-1 {
		y++
		if treeGrid[i][y] >= treeGrid[i][j] {
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
