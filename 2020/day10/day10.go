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
	fmt.Println(diffs[0] * diffs[2])
}
