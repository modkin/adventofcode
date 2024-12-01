package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("2024/day1/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int
	for scanner.Scan() {
		//lines = append(lines, scanner.Text())
		split := strings.Fields(scanner.Text())
		left = append(left, utils.ToInt(split[0]))
		right = append(right, utils.ToInt(split[1]))
	}
	slices.Sort(left)
	slices.Sort(right)

	one := 0
	for i, i2 := range left {
		one += utils.IntAbs(right[i] - i2)
	}

	two := 0
	for _, i2 := range left {
		two += i2 * utils.CountInSlice(right, i2)
	}
	fmt.Println("Day 1.1:", one)
	fmt.Println("Day 2.1:", two)

}
