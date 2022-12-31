package main

import (
	"bufio"
	"fmt"
	"os"
)

func uniqueElementsInQueue(queue []rune) bool {
	queueMap := make(map[rune]bool)

	for _, item := range queue {
		if _, ok := queueMap[item]; ok {
			return false
		}
		queueMap[item] = true
	}
	return true
}

func enqueue(queue []rune, char rune) ([]rune, bool) {
	if len(queue) == 14 {
		if uniqueElementsInQueue(queue) {
			return queue, true
		}
		queue = dequeue(queue)
	}
	queue = append(queue, char)
	return queue, false
}

func dequeue(queue []rune) []rune {
	return queue[1:]
}

func main() {
	var startOfPacketMarkerPos int
	var txt string
	var queue []rune
	var unique bool

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt = scanner.Text()
	}

	for idx, ch := range txt {
		queue, unique = enqueue(queue, ch)
		if unique {
			startOfPacketMarkerPos = idx
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(startOfPacketMarkerPos)
}
