package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("2021/day1/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	count1 := 0
	count2 := 0
	measures := make([]int, 0)
	for scanner.Scan() {
		measures = append(measures, utils.ToInt(scanner.Text()))
	}
	for i := 1; i < len(measures); i++ {
		if measures[i] > measures[i-1] {
			count1++
		}
	}

	for i := 1; i < len(measures)-2; i++ {
		prev := measures[i-1] + measures[i] + measures[i+1]
		next := measures[i] + measures[i+1] + measures[i+2]
		if next > prev {
			count2++
		}
	}
	fmt.Println("Day 1.1:", count1)
	fmt.Println("Day 1.2:", count2)
}
