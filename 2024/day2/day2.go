package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isReportSafe(input []string) bool {
	inc := true
	dec := true
	dist := true

	for i, i2 := range input {
		if i == 0 {
			continue
		}
		if utils.ToInt(i2) < utils.ToInt(input[i-1]) {
			inc = false
		}
		if utils.ToInt(i2) > utils.ToInt(input[i-1]) {
			dec = false
		}
		if utils.IntAbs(utils.ToInt(i2)-utils.ToInt(input[i-1])) > 3 {
			dist = false
		}
		if utils.IntAbs(utils.ToInt(i2)-utils.ToInt(input[i-1])) < 1 {
			dist = false
		}
	}
	return (inc || dec) && dist
}

func removeI(in []string, i int) []string {
	cop := utils.CopyStringSlice(in)
	return append(cop[:i], cop[i+1:]...)
}

func main() {
	file, err := os.Open("2024/day2/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	safe1 := 0
	safe2 := 0

	scanner := bufio.NewScanner(file)
outer:
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		if isReportSafe(split) {
			safe1++
			safe2++
			continue
		} else {
			for i := range split {
				if isReportSafe(removeI(split, i)) {
					safe2++
					continue outer

				}
			}
		}

	}

	fmt.Println("Day 2.1:", safe1)
	fmt.Println("Day 2.2:", safe2)
}
