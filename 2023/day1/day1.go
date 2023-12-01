package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func firstAndLastStringInt(inputString string) (first string, last string, firstPos int, lastPos int) {
	numberStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	firstPos = len(inputString) - 1
	lastPos = 0
	var front, back string
	for i := 0; i < len(inputString); i++ {
		front += string(inputString[i])
		back = string(inputString[len(inputString)-i-1]) + back
		for j, numberString := range numberStrings {
			if strings.Contains(front, numberString) && first == "" {
				first = strconv.Itoa(j + 1)
				firstPos = i
			}
			if strings.Contains(back, numberString) && last == "" {
				last = strconv.Itoa(j + 1)
				lastPos = len(inputString) - i
			}
		}
	}
	return
}

func firstAndLastIntInString(input string) (first string, last string, firstPos int, lastPos int) {
	first = "0"
	last = "0"
	firstPos = len(input) - 1
	lastPos = 0
	for pos, char := range input {
		if unicode.IsDigit(char) {
			if first == "0" {
				first = string(char)
				firstPos = pos
			}
			last = string(char)
			lastPos = pos
		}
	}
	return
}

func main() {
	file, err := os.Open("2023/day1/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	var firstInt, lastInt string
	var firstPos, lastPos int
	var valuesPart1 []int
	var valuesPart2 []int

	for _, line := range lines {
		firstInt, lastInt, firstPos, lastPos = firstAndLastIntInString(line)
		firstStringInt, lastStringInt, firstStringPos, lastStringPos := firstAndLastStringInt(line)
		valuesPart1 = append(valuesPart1, utils.ToInt(firstInt+lastInt))
		if firstStringPos < firstPos {
			firstInt = firstStringInt
		}
		if lastStringPos > lastPos {
			lastInt = lastStringInt
		}
		valuesPart2 = append(valuesPart2, utils.ToInt(firstInt+lastInt))
	}
	valuesSum1 := utils.SumSlice(valuesPart1)
	fmt.Println("Day 1.1:", valuesSum1)
	valuesSum2 := utils.SumSlice(valuesPart2)
	fmt.Println("Day 1.2:", valuesSum2)
}
