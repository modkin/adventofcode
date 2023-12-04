package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2023/day4/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	totalPoints := 0

	multiplier := make([]int, len(lines))
	for i, _ := range multiplier {
		multiplier[i] = 1
	}

	for idx, line := range lines {
		offset := 1
		split := strings.Split(line, ":")
		winMine := strings.Split(split[1], "|")
		var win []string
		CardPoints := 0
		for _, s := range strings.Fields(strings.TrimSpace(winMine[0])) {
			win = append(win, s)
		}
		for _, s := range strings.Split(winMine[1], " ") {
			for _, s2 := range win {
				if s == s2 {
					multiplier[idx+offset] += multiplier[idx]
					offset++
					if CardPoints == 0 {
						CardPoints = 1
					} else {
						CardPoints *= 2
					}
				}
			}
		}
		totalPoints += CardPoints
	}

	fmt.Println("Day 4.1:", totalPoints)
	sum := utils.SumSlice(multiplier)
	fmt.Println("Day 4.2:", sum)
}
