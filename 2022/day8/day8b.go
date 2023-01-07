package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getScenicScore(treeGrid [][]int, row, col int) int {
	i := row
	scenicScoreUp := 0
	for i > 0 {
		i--
		if treeGrid[i][col] < treeGrid[row][col] {
			scenicScoreUp++
		} else {
			scenicScoreUp++
			break
		}
	}

	scenicScoreDown := 0
	i = row
	for i < len(treeGrid)-1 {
		i++
		if treeGrid[i][col] < treeGrid[row][col] {
			scenicScoreDown++
		} else {
			scenicScoreDown++
			break
		}
	}

	j := col
	scenicScoreLeft := 0
	for j > 0 {
		j--
		if treeGrid[row][j] < treeGrid[row][col] {
			scenicScoreLeft++
		} else {
			scenicScoreLeft++
			break
		}
	}

	scenicScoreRight := 0
	j = col
	for j < len(treeGrid[0])-1 {
		j++
		if treeGrid[row][j] < treeGrid[row][col] {
			scenicScoreRight++
		} else {
			scenicScoreRight++
			break
		}
	}

	return scenicScoreUp * scenicScoreDown * scenicScoreLeft * scenicScoreRight
}

func main() {
	var scenicScore int
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

	var computedScenicScore int
	for i := 1; i < len(treeGrid)-1; i++ {
		for j := 1; j < len(treeGrid[i])-1; j++ {
			computedScenicScore = getScenicScore(treeGrid, i, j)
			if computedScenicScore >= scenicScore {
				scenicScore = computedScenicScore
			}
		}
	}
	fmt.Println(scenicScore)
}
