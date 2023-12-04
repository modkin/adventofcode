package main

import (
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

	//var lines []string

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	totalPoints := 0

	//var multiplier []int
	idx := 1
	for _, line := range lines {
		split := strings.Split(line, ":")
		winMine := strings.Split(split[1], "|")
		var win []string
		//var mine []string
		CardPoints := 0
		//multiplierIdx := idx
		for _, s := range strings.Fields(strings.TrimSpace(winMine[1])) {
			win = append(win, s)
		}
		for _, s := range strings.Split(winMine[0], " ") {
			for _, s2 := range win {
				if s == s2 {

					if CardPoints == 0 {
						CardPoints = 1
					} else {
						CardPoints *= 2
					}
				}
			}
		}
		totalPoints += CardPoints
		idx++
	}

	fmt.Println(totalPoints)
	//sum := utils.SumSlice(lines)
	//fmt.Println("Day 3.2:", sum)
}
