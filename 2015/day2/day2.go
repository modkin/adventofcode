package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2015/day2/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		dimensions := strings.Split(scanner.Text(), "x")
		x, y, z := utils.ToInt(dimensions[0]), utils.ToInt(dimensions[1]), utils.ToInt(dimensions[2])
		XY := x * y
		YZ := y * z
		XZ := x * z
		total += 2*XY + 2*YZ + 2*XZ
		if XY <= YZ && XY <= XZ {
			total += XY
		} else if YZ <= XZ {
			total += YZ
		} else {
			total += XZ
		}
	}
	fmt.Println("Task 2.1: ", total)
}
