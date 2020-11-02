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
	totalPaper := 0
	totalRibbon := 0
	for scanner.Scan() {
		dimensions := strings.Split(scanner.Text(), "x")
		x, y, z := utils.ToInt(dimensions[0]), utils.ToInt(dimensions[1]), utils.ToInt(dimensions[2])
		XY := x * y
		YZ := y * z
		XZ := x * z
		totalPaper += 2*XY + 2*YZ + 2*XZ
		if XY <= YZ && XY <= XZ {
			totalPaper += XY
		} else if YZ <= XZ {
			totalPaper += YZ
		} else {
			totalPaper += XZ
		}

		totalRibbon += x * y * z
		if x > y && x > z {
			totalRibbon += 2*y + 2*z
		} else if y > z {
			totalRibbon += 2*x + 2*z
		} else {
			totalRibbon += 2*x + 2*y
		}
	}
	fmt.Println("Task 2.1: ", totalPaper)
	fmt.Println("Task 2.2: ", totalRibbon)
}
