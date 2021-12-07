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
	min := 1000000
	max := 0
	for _, crab := range crabsPos {
		if crab < min {
			min = crab
		}
		if crab > max {
			max = crab
		}
	}
	bestAlign := 10000000
	bestPos := 0
	for i := min; i <= max; i++ {
		tmp := 0.0
		for _, crab := range crabsPos {
			tmp += math.Abs(float64(i - crab))
		}
		fmt.Println(i, tmp)
		if int(tmp) < bestAlign {
			bestAlign = int(tmp)
			bestPos = i
		}
	}
	fmt.Println(bestPos)
	fmt.Println(bestAlign)
}
