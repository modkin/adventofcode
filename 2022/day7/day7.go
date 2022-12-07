package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func addRecursive(dirContents map[string][]string, dirSizes map[string]int, path []string) {
	currentDir := strings.Join(path, "-")
	for _, i := range dirContents[currentDir] {
		subdir := append(path, i)
		for i := range path {
			dirSizes[strings.Join(path[:i+1], "-")] += dirSizes[strings.Join(subdir, "-")]
		}
		newPath := utils.CopyStringSlice(path)
		newPath = append(newPath, i)
		addRecursive(dirContents, dirSizes, newPath)
	}
}

func main() {

	file, err := os.Open("2022/day7/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	commands := make([][]string, 0)
	dirSizes := make(map[string]int)
	dirContents := make(map[string][]string)
	//partentDir := make(map[string]string)
	//foo := make(map[string]int)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		commands = append(commands, split)
	}

	var currentDir []string
outer:
	for i := 0; i < len(commands); {
		if commands[i][0] == "$" {
			if commands[i][1] == "cd" {
				if commands[i][2] == ".." {
					currentDir = currentDir[:len(currentDir)-1]
				} else {
					//partentDir[commands[i][2]] = currentDir
					currentDir = append(currentDir, commands[i][2])
				}
			} else if commands[i][1] == "ls" {
				i++
				for commands[i][0] != "$" {
					if commands[i][0] == "dir" {
						dirContents[strings.Join(currentDir, "-")] = append(dirContents[strings.Join(currentDir, "-")], commands[i][1])
					} else {
						dirSizes[strings.Join(currentDir, "-")] += utils.ToInt(commands[i][0])
					}
					i++
					if i >= len(commands) {
						break outer
					}
				}
				continue
			}
		}
		i++

	}
	fmt.Println(dirContents)
	fmt.Println(dirSizes)
	addRecursive(dirContents, dirSizes, []string{"/"})

	total := 0
	fmt.Println(dirSizes)
	for _, i := range dirSizes {
		if i <= 100000 {
			total += i
		}
	}

	fmt.Println("Day 6.1: ", total)

	totalSize := 70000000
	required := 30000000
	currentFree := totalSize - dirSizes["/"]
	fmt.Println(currentFree)
	SizeList := make([]int, 0)
	for _, i := range dirSizes {
		SizeList = append(SizeList, i)
	}
	sort.Ints(SizeList)
	for _, i := range SizeList {
		if currentFree+i >= required {
			fmt.Println(i)
			break
		}
	}
}
