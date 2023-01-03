package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxAvailableSpace = 70000000
	updateSpace       = 30000000
)

var (
	rootDir        = Directory{name: "/"}
	currentDir     = &rootDir
	minNeededSpace = 0
)

type File struct {
	name string
	size int
}

type Directory struct {
	name      string
	totalSize int
	files     []File
	dirs      []*Directory
	parent    *Directory
}

func propagateSizeToRoot(size int) {
	parent := currentDir.parent
	for parent != nil {
		parent.totalSize += size
		parent = parent.parent
	}
}

func findDirToDelete(dir *Directory, neededSpace int) {
	if dir.totalSize >= neededSpace && minNeededSpace > dir.totalSize {
		minNeededSpace = dir.totalSize
	}
	for _, d := range dir.dirs {
		findDirToDelete(d, neededSpace)
	}
	return
}

func getSmallestDir() {
	neededSpace := updateSpace - (maxAvailableSpace - rootDir.totalSize)
	minNeededSpace = rootDir.totalSize
	currentD := &rootDir
	findDirToDelete(currentD, neededSpace)
}

func processInput(str string) {
	if strings.HasPrefix(str, "$") {
		command := strings.Split(str, " ")
		if command[1] == "cd" {
			if command[2] == "/" {
				currentDir = &rootDir
				return
			} else if command[2] == ".." {
				currentDir = currentDir.parent
				return
			} else {
				for _, dir := range currentDir.dirs {
					if dir.name == command[2] {
						currentDir = dir
						return
					}
				}
			}
		} else if command[1] == "ls" {
			return
		}
	}
	line := strings.Split(str, " ")
	if line[0] == "dir" {
		newDir := Directory{name: line[1], parent: currentDir}
		currentDir.dirs = append(currentDir.dirs, &newDir)
		return
	}
	fileSize, err := strconv.Atoi(line[0])
	if err != nil {
		panic(err)
	}
	fileName := line[1]
	newFile := File{name: fileName, size: fileSize}
	currentDir.totalSize += fileSize
	currentDir.files = append(currentDir.files, newFile)
	propagateSizeToRoot(fileSize)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		processInput(txt)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	getSmallestDir()
	fmt.Println(minNeededSpace)
}
