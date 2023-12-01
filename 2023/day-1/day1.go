package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2023/day-1/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	var allInts []int
	var intPos []int
outer:
	for _, val := range lines {
		var result []string
		ints := strings.Split(val, "")
		for pos, i := range ints {
			if res, err := strconv.Atoi(string(i)); err == nil {
				result = append(result, strconv.Itoa(res))
				intPos = append(intPos, pos)
			}
		}
		if len(result) == 1 {
			allInts = append(allInts, utils.ToInt(result[0]+result[0]))
			continue outer
		} else {
			allInts = append(allInts, utils.ToInt(result[0]+result[len(result)-1]))
			continue outer
		}
	}

	day1 := utils.SumSlice(allInts)

	fmt.Println(allInts)
	fmt.Println("Day 1.1", day1)
}
