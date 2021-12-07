package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2021/day7/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	crabsPos := make([]int, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, s := range line {
			crabsPos = append(crabsPos, utils.ToInt(s))
		}
	}
	bestAlign := math.MaxInt64
	bestAlign1 := math.MaxInt64

	for _, i := range crabsPos {
		tmp := 0
		for _, crab := range crabsPos {
			tmp += utils.IntAbs(i - crab)
		}
		if tmp < bestAlign1 {
			bestAlign1 = tmp
		}
		totalFuel := 0
		for _, crab := range crabsPos {
			tmp = 0
			tmp += utils.IntAbs(i - crab)
			for j := 1; j <= tmp; j++ {
				totalFuel += j
			}
		}
		if totalFuel < bestAlign {
			bestAlign = totalFuel
		}
	}
	fmt.Println("Day 7.1:", bestAlign1)
	fmt.Println("Day 7.2:", bestAlign)
}
