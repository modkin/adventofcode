package utils

import (
	"math"
	"os"
	"strconv"
	"strings"
)

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
