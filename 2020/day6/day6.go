package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2020/day6/input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	currentGroup := make(map[string]bool)
	currentGroup2 := make(map[string]bool)
	total, total2 := 0, 0
	first := true
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		if len(line) == 0 {
			total += len(currentGroup)
			total2 += len(currentGroup2)
			currentGroup = make(map[string]bool)
			currentGroup2 = make(map[string]bool)
			first = true
		} else {
			for _, answer := range line {
				currentGroup[answer] = true
			}
			if first == true {
				for _, answer := range line {
					currentGroup2[answer] = true
				}
				first = false
			} else {
				tmp := make(map[string]bool)
				for _, answer := range line {
					if currentGroup2[answer] {
						tmp[answer] = true
					}
				}
				currentGroup2 = tmp
			}
		}
	}
	fmt.Println("Task 6.1:", total+len(currentGroup))
	fmt.Println("Task 6.2:", total2+len(currentGroup2))
}
