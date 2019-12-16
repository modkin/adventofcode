package main

import (
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
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
	inputSplit := strings.Split(string(inputString), "")
	input := make([]int, len(inputSplit))
	for i, char := range inputSplit {
		input[i] = utils.ToInt(char)
	}

	nextInput := make([]int, len(input))
	for phase := 0; phase < 100; phase++ {
		for digit := 0; digit < len(input); digit++ {
			factor := genPattern(digit + 1)
			result := 0
			for i := 0; i < len(input); i++ {
				result += input[i] * factor()
			}
			nextInput[digit] = utils.IntAbs(result % 10)
		}
		input = nextInput
	}
	fmt.Println(input[:8])

}
