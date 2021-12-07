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
	min := math.MaxInt64
	max := 0
	for _, crab := range crabsPos {
		if crab < min {
			min = crab
		}
		if crab > max {
			max = crab
		}
	}
	bestAlign := math.MaxInt64
	bestPos := 0
	for i := min; i <= max; i++ {
		totalFuel := 0
		for _, crab := range crabsPos {
			tmp := 0
			tmp += utils.IntAbs(i - crab)
			for j := 1; j <= tmp; j++ {
				totalFuel += j
			}
		}
		if totalFuel < bestAlign {
			bestAlign = totalFuel
			bestPos = i
		}
	}
	fmt.Println(bestPos)
	fmt.Println(bestAlign)
}
