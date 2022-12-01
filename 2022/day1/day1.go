package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	file, err := os.Open("2022/day1/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	measures := make([]int, 0)
	oneElf := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			measures = append(measures, oneElf)
			oneElf = 0
		} else {
			oneElf += utils.ToInt(scanner.Text())
		}

	}
	sort.Ints(measures)

	total := 0
	for i := 1; i < 4; i++ {
		total += measures[len(measures)-i]
	}

	fmt.Println("Day 1.1:", measures[len(measures)-1])
	fmt.Println("Day 1.2:", total)

}
