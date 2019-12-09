package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ToInt64(str string) int64 {
	ret, err := strconv.ParseInt(strings.TrimSpace(str), 10, 64)
	if err != nil {
		panic(err)
	}
	return int64(ret)
}

func ToInt(str string) int {
	ret, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		panic(err)
	}
	return ret
}

func IntAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func SumSlice(slice []int) int {
	sum := 0
	for _, element := range slice {
		sum += element
	}
	return sum
}

func ReverseSlice(a []string) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

///split integer into integer slice of the digits
func SplitInt(input int) []int {
	count := len(fmt.Sprint(input))
	output := make([]int, count)
	for i := 0; i < count; i++ {
		output[i] = input / int(math.Pow(10, float64(count-1-i))) % 10
	}
	return output
}
