package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2023/day9/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var history [][]int
	for scanner.Scan() {
		//line := append(lines, scanner.Text())
		newLine := make([]int, 0)
		for _, s := range strings.Fields(scanner.Text()) {
			newLine = append(newLine, utils.ToInt(s))
		}
		history = append(history, newLine)

	}
	//fmt.Println(history)
	getDiffs := func(input []int) []int {
		var diffs []int
		for i := 0; i < len(input)-1; i++ {
			diffs = append(diffs, input[i+1]-input[i])
		}
		return diffs
	}
	totalSum := 0
	totalSum2 := 0
	var offsets []int
	for i := 0; i < len(history); i++ {
		var diffs [][]int
		diff := utils.CopyIntSlice(history[i])
		diffs = append(diffs, diff)
		for {
			diff = utils.CopyIntSlice(getDiffs(diff))
			diffs = append(diffs, diff)
			allZero := true
			for _, i2 := range diff {
				if i2 != 0 {
					allZero = false
				}
			}
			if allZero {
				break
			}
		}
		offset := 0
		offset2 := 0
		for j := len(diffs) - 1; j >= 0; j-- {
			offset += diffs[j][len(diffs[j])-1]
			offset2 = diffs[j][0] - offset2
		}
		totalSum += offset
		totalSum2 += offset2
		offsets = append(offsets, offset)
		//fmt.Print(offset, " ")
	}

	fmt.Println("Day 9.1:", totalSum)
	fmt.Println("Day 9.2:", totalSum2)
}
