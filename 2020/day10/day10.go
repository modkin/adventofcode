package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"sort"
)

func findNext(index int, jolt int, adapters []int, target int) int {
	nOA := 0
	for i := 1; i <= 3; i++ {
		if jolt+3 == target {
			nOA++
			break
		}
		tmp := index + i
		if tmp > len(adapters)-1 {
			continue
		}
		nextJolt := adapters[tmp]
		if nextJolt <= jolt+3 {
			if i == 3 {
				fmt.Println("bla")
			}
			nOA += findNext(tmp, nextJolt, adapters, target)
		}

	}
	return nOA
}

func main() {

	adapters := make([]int, 0)
	scanner := bufio.NewScanner(utils.OpenFile("2020/day10/testinput"))
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
	fmt.Println(diffs)
	fmt.Println(diffs[0] * diffs[2])

	foo := []int{0}
	adapters = append(foo, adapters...)
	targetJolt := adapters[len(adapters)-1] + 3
	numberOfArrangements := findNext(0, 0, adapters, targetJolt)
	//for {
	//
	//	for i:= 3; i >= 1; i-- {
	//
	//		if adapters[tmp]
	//	}
	//}
	fmt.Println(numberOfArrangements)
}
