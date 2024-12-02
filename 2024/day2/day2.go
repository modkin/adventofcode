package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func increase(input []string) bool {
	inc := true
	dec := true
	dist := true

	var wrong1 []int
	var wrong2 []int
	for i, i2 := range input {
		if i == 0 {
			continue
		}
		if utils.ToInt(i2) < utils.ToInt(input[i-1]) {
			inc = false
			wrong1 = append(wrong1, i)
		}
		if utils.ToInt(i2) > utils.ToInt(input[i-1]) {
			dec = false
			wrong1 = append(wrong1, i)
		}
		if utils.IntAbs(utils.ToInt(i2)-utils.ToInt(input[i-1])) > 3 {
			dist = false
			wrong2 = append(wrong2, i)
		}
		if utils.IntAbs(utils.ToInt(i2)-utils.ToInt(input[i-1])) < 1 {
			dist = false
			wrong2 = append(wrong2, i)
		}
	}
	return (inc || dec) && dist
}

func main() {
	file, err := os.Open("2024/day2/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//var one []string
	safe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//lines = append(lines, scanner.Text())
		split := strings.Fields(scanner.Text())
		if increase(split) {
			safe++
			fmt.Println(split)
		}

	}

	fmt.Println("Day 2.1:", safe)
	//fmt.Println("Day 2.2:", two)
}
