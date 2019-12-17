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
	fmt.Println(idx)
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

	for i := 0; i < 8; i++ {
		result := int64(0)
		for j, elem := range input[idxStart+i:] {
			result += matrix[j] * int64(elem)
		}
		fmt.Println(result)
	}

	//
	//fmt.Println(idxStart)
	//nextInput := make([]int, len(input))
	//for phase := 0; phase < 100; phase++ {
	//	t0 := time.Now()
	//	for digit := idxStart; digit < len(input); digit++ {
	//		//factor := genPattern(digit + 1)
	//		result := 0
	//		for i := idxStart; i < len(input); i++ {
	//			if i >= digit {
	//				result += input[i] * idxStart
	//			}
	//		}
	//		//result *= 10_000
	//		nextInput[digit] = utils.IntAbs(result % 10)
	//	}
	//	input = nextInput
	//	fmt.Println(phase, " took: ", time.Now().Sub(t0))
	//}

	//fmt.Println(input[idxStart:idxStart+8])

}
