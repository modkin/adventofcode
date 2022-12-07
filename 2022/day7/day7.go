package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	file, err := os.Open("2022/day7/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	commands := make([][]string, 0)
	dirSizes := make(map[string]int)

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
					currentDir = append(currentDir, commands[i][2])
				}
			} else if commands[i][1] == "ls" {
				i++
				for commands[i][0] != "$" {
					if commands[i][0] != "dir" {
						for i2 := range currentDir {
							dirSizes[strings.Join(currentDir[:i2+1], "-")] += utils.ToInt(commands[i][0])
						}
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

	total := 0
	for _, i := range dirSizes {
		if i <= 100000 {
			total += i
		}
	}

	fmt.Println("Day 7.1:", total)
	if total != 1783610 {
		panic(err)
	}

	required := 30000000
	currentFree := 70000000 - dirSizes["/"]
	SizeList := make([]int, 0)
	for _, i := range dirSizes {
		SizeList = append(SizeList, i)
	}
	sort.Ints(SizeList)
	for _, i := range SizeList {
		if currentFree+i >= required {
			fmt.Println("Day 7.2:", i)
			if i != 4370655 {
				panic(err)
			}
			break
		}
	}
}
