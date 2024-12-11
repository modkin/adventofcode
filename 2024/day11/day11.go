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

func do5Blinks(stones []int) (output []int) {
	output = utils.CopyIntSlice(stones)
	for i := 0; i < 5; i++ {
		output = blink(output)
	}
	return output
}
func doHashedBlinks(stones []int, hash map[int][]int, totalSteps int) (output []int) {

	for steps := 0; steps < totalSteps; steps++ {
		var newStone2 []int
		for _, i2 := range stones {
			if val, ok := hash[i2]; ok {
				newStone2 = append(newStone2, val...)
			} else {
				tmp5 := do5Blinks([]int{i2})
				newStone2 = append(newStone2, tmp5...)
				hash[i2] = tmp5
			}
		}
		stones = newStone2
		//fmt.Println(steps, len(stones))
	}
	return stones
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day11/input")

	stones := make([]int, 0)
	for _, line := range lines {
		split := strings.Split(line, " ")
		for _, s := range split {
			stones = append(stones, utils.ToInt(s))
		}
	}

	hash5 := make(map[int][]int)
	stones = doHashedBlinks(stones, hash5, 5)
	fmt.Println("Day 11.1:", len(stones))

	stones = doHashedBlinks(stones, hash5, 3)

	hash35 := make(map[int]int)
	sum := 0
	for _, i2 := range stones {
		if val, ok := hash35[i2]; ok {
			sum += val
		} else {
			tmp35 := doHashedBlinks([]int{i2}, hash5, 7)
			hash35[i2] = len(tmp35)
			sum += len(tmp35)
		}
	}
	fmt.Println("Day 11.2:", sum)

}
