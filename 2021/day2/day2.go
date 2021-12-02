package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2021/day2/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	pos1, pos2 := 0, 0
	height1, height2 := 0, 0
	aim2 := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir, amount := line[0], utils.ToInt(line[1])
		if dir == "forward" {
			pos1 += amount
			pos2 += amount
			height2 += aim2 * amount
		} else if dir == "up" {
			height1 -= amount
			aim2 -= amount
		} else if dir == "down" {
			height1 += amount
			aim2 += amount
		}
	}
	fmt.Println("Day 2.1:", pos1*height1)
	fmt.Println("Day 2.2:", pos2*height2)
}
