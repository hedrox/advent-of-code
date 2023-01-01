package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getScenicScore(treeGrid [][]int, i, j int) int {
	x := i
	scenicScoreUp := 0
	for x > 0 {
		x--
		if treeGrid[x][j] < treeGrid[i][j] {
			scenicScoreUp++
		} else {
			scenicScoreUp++
			break
		}
	}

	scenicScoreDown := 0
	x = i
	for x < len(treeGrid)-1 {
		x++
		if treeGrid[x][j] < treeGrid[i][j] {
			scenicScoreDown++
		} else {
			scenicScoreDown++
			break
		}
	}

	y := j
	scenicScoreLeft := 0
	for y > 0 {
		y--
		if treeGrid[i][y] < treeGrid[i][j] {
			scenicScoreLeft++
		} else {
			scenicScoreLeft++
			break
		}
	}

	scenicScoreRight := 0
	y = j
	for y < len(treeGrid[0])-1 {
		y++
		if treeGrid[i][y] < treeGrid[i][j] {
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
