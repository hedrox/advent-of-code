package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var currCalories int
	var totalCalories int

	h := &IntHeap{}
	heap.Init(h)

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		// Empty line
		if len(txt) == 0 {
			heap.Push(h, currCalories)
			currCalories = 0
		} else {
			val, _ := strconv.Atoi(txt)
			currCalories += val
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// There are still some unprocessed calories
	if currCalories != 0 {
		heap.Push(h, currCalories)
	}

	mostCalories := heap.Pop(h).(int)
	totalCalories += mostCalories
	for i := 0; i < 2; i++ {
		totalCalories += heap.Pop(h).(int)
	}
	fmt.Printf("Most calories carried by an elf: %v\n", mostCalories)
	fmt.Printf("Total number of calories carried by top 3 elfs: %v\n", totalCalories)
}
