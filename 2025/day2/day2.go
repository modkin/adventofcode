package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ranges := strings.Split(utils.ReadFileIntoLines("2025/day2/input")[0], ",")

	invalidSum := 0
	for _, r := range ranges {
		startStr, endStr, _ := strings.Cut(r, "-")
		startInt := utils.ToInt(startStr)
		endInt := utils.ToInt(endStr)
		if len(startStr) != len(endStr) {
			//fmt.Println("that was not expected")
			//return
		}

		halfStart := startStr[0 : len(startStr)/2]
		if halfStart == "" {
			halfStart = startStr[0:1]
		}
		halfEnd := endStr[0 : (len(endStr)+1)/2]

		fmt.Println("range: ", startStr, "-", endStr)
		for i := utils.ToInt(halfStart); i <= utils.ToInt(halfEnd); i++ {
			num := utils.ToInt(strconv.Itoa(i) + strconv.Itoa(i))
			if startInt <= num && num <= endInt {
				fmt.Println("invalid: ", startStr, "-", endStr, " ", num)
				invalidSum += num
			}
		}
	}
	fmt.Println("Day 2.1: ", invalidSum)
}
