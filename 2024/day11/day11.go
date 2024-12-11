package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SplitInt(input int) []int {
	count := len(fmt.Sprint(input))
	output := make([]int, count)
	for i := 0; i < count; i++ {
		output[i] = input / int(math.Pow(10, float64(count-1-i))) % 10
	}
	return output
}

func intSlicetoInt(input []int) int {
	str := ""
	for _, i2 := range input {
		str += strconv.Itoa(i2)
	}
	return utils.ToInt(str)
}

func blink(stones []int) (output []int) {
	stonePos := 0
	for stonePos < len(stones) {
		stone := stones[stonePos]
		digits := SplitInt(stones[stonePos])
		if stone == 0 {
			stones[stonePos] = 1
		} else if len(digits)%2 == 0 {
			stones[stonePos] = intSlicetoInt(digits[0 : len(digits)/2])
			copyI := utils.CopyIntSlice(stones[stonePos+1:])
			step1 := append(stones[:stonePos+1], intSlicetoInt(digits[len(digits)/2:len(digits)]))
			stones = append(step1, copyI...)
			stonePos++
		} else {
			stones[stonePos] *= 2024
		}
		stonePos++
	}
	return stones
}

func countMap(input map[int]int) (sum int) {
	for _, count := range input {
		sum += count
	}
	return
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day11/input")

	stoneCount := make(map[int]int)
	for _, line := range lines {
		split := strings.Split(line, " ")
		for _, s := range split {
			stoneCount[utils.ToInt(s)]++
		}
	}

	for i := 0; i < 75; i++ {
		if i == 25 {
			fmt.Println("Day 11.1:", countMap(stoneCount))
		}
		newStoneCount := make(map[int]int)
		for stone, count := range stoneCount {
			tmp := blink([]int{stone})
			for _, st := range tmp {
				newStoneCount[st] += count
			}
		}
		stoneCount = newStoneCount
	}
	fmt.Println("Day 11.2:", countMap(stoneCount))
}
