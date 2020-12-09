package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
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

func sumMinMax(list []int) int {
	min, max := math.MaxInt32, 0
	for _, elem := range list {
		if elem > max {
			max = elem
		}
		if elem < min {
			min = elem
		}
	}
	return min + max
}

func main() {
	previous := make([]int, 0)
	allNumbers := make([]int, 0)
	scanner := bufio.NewScanner(utils.OpenFile("2020/day9/input"))
	count, target := 0, 0
	for scanner.Scan() {
		allNumbers = append(allNumbers, utils.ToInt(scanner.Text()))
		if count <= 25 {
			previous = append(previous, utils.ToInt(scanner.Text()))
			count++
		} else {
			current := utils.ToInt(scanner.Text())
			if !checkNumber(previous, current) {
				target = current
				break
			}
			previous = append(previous[1:], current)
		}
	}
	fmt.Println("Task 9.1:", target)
	currentStart := 0
outer:
	for {
		sum := allNumbers[currentStart]
		for i := 1; i < target; i++ {
			sum += allNumbers[currentStart+i]
			if sum == target {
				fmt.Println("Task 9.2:", sumMinMax(allNumbers[currentStart:currentStart+i+1]))
				break outer
			} else if sum > target {
				currentStart++
				break
			}
		}
	}
}
