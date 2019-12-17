package main

import (
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func genPattern(pos int) func() int {
	initialPattern := []int{0, 1, 0, -1}
	counter := 0
	return func() int {
		counter++
		idx := (counter / pos) % 4
		return initialPattern[idx]
	}
}

func main() {
	inputByte, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := strings.TrimSuffix(string(inputByte), "\n")
	inputString = strings.Repeat(inputString, 10_000)
	inputSplit := strings.Split(string(inputString), "")
	input := make([]int, len(inputSplit))
	for i, char := range inputSplit {
		input[i] = utils.ToInt(char)
	}

	idx := input[:7]
	idxStart := 0
	for i := 0; i < 7; i++ {
		idxStart += idx[i] * int(math.Pow(10, float64(6-i)))
	}

	matrix := make([]int64, len(input[idxStart:]))

	for i := 0; i < len(matrix); i++ {
		matrix[i] = 1
	}
	for phase := 0; phase < 99; phase++ {
		for i := 1; i < len(matrix); i++ {
			matrix[i] = (matrix[i] + matrix[i-1]) % 10
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i] < 0 {
			panic("smaller 0")
		}
	}

	var task2 []int64
	for i := 0; i < 8; i++ {
		result := int64(0)
		for j, elem := range input[idxStart+i:] {
			result += matrix[j] * int64(elem)
		}
		task2 = append(task2, result%10)
	}
	task2String := strings.Trim(strings.Replace(fmt.Sprint(task2), " ", "", -1), "[]")

	nextInput := make([]int, len(input))
	for phase := 0; phase < 100; phase++ {
		for digit := 0; digit < len(input)/10_000; digit++ {
			factor := genPattern(digit + 1)
			result := 0
			for i := 0; i < len(input)/10_000; i++ {
				result += input[i] * factor()
			}
			nextInput[digit] = utils.IntAbs(result % 10)
		}
		input = nextInput
	}

	task1 := strings.Trim(strings.Replace(fmt.Sprint(input[:8]), " ", "", -1), "[]")
	fmt.Println("Task 16.1: ", task1)
	fmt.Println("Task 16.2: ", task2String)

}
