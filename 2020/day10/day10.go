package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"sort"
)

func main() {

	adapters := make([]int, 0)
	scanner := bufio.NewScanner(utils.OpenFile("2020/day10/input"))
	for scanner.Scan() {
		adapters = append(adapters, utils.ToInt(scanner.Text()))
	}
	sort.Ints(adapters)
	diffs := []int{0, 0, 0}
	currentJolt := 0
	for _, elem := range adapters {
		diffs[elem-currentJolt-1]++
		currentJolt = elem
	}
	diffs[2]++
	fmt.Println("Task 10.1:", diffs[0]*diffs[2])

	adapters = append([]int{0}, adapters...)
	targetJolt := adapters[len(adapters)-1] + 3
	adapters = append(adapters, targetJolt)
	numberOfArrangements, index := 1, 0
	currentJolt = 0

	for {
		offset := 0
		for {
			offset++
			if index+offset == len(adapters)-1 {
				break
			}
			if adapters[index+offset] != currentJolt+offset {
				break
			}
		}
		if offset == 1 {
			index++
			currentJolt = adapters[index]
		} else if offset == 2 {
			index += offset
			currentJolt = adapters[index]
			numberOfArrangements *= 1
		} else if offset == 3 {
			index += offset
			currentJolt = adapters[index]
			numberOfArrangements *= 2
		} else if offset == 4 {
			index += offset
			currentJolt = adapters[index]
			numberOfArrangements *= 4
		} else if offset == 5 {
			index += offset
			currentJolt = adapters[index]
			numberOfArrangements *= 7
		}
		if adapters[index]+3 == targetJolt {
			break
		}
	}
	fmt.Println("Task 10.2:", numberOfArrangements)
}
