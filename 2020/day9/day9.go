package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
)

func checkNumber(previous []int, current int) bool {
	for _, x := range previous {
		for _, y := range previous {
			if x+y == current {
				return true
			}
		}
	}
	return false
}

func main() {
	previous := make([]int, 0)
	scanner := bufio.NewScanner(utils.OpenFile("2020/day9/input"))
	count := 0
	for scanner.Scan() {
		if count <= 25 {
			previous = append(previous, utils.ToInt(scanner.Text()))
			count++
		} else {
			current := utils.ToInt(scanner.Text())
			if !checkNumber(previous, current) {
				fmt.Println(current)
				break
			}
			previous = append(previous[1:], current)
		}

	}
}
