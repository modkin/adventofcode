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
	pos := 0
	height := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir, amount := line[0], utils.ToInt(line[1])
		if dir == "forward" {
			pos += amount
		}
		if dir == "up" {
			height += amount
		}
		if dir == "down" {
			height -= amount
		}
	}
	fmt.Println(pos * height)
}
