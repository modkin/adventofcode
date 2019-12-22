package main

import (
	"adventofcode/utils"
	"fmt"
	"testing"
)

func TestTask1(t *testing.T) {
	testStack := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testStack = dealWithInc(7, testStack)
	testStack = utils.ReverseIntSlice(testStack)
	testStack = utils.ReverseIntSlice(testStack)
	fmt.Println(testStack)
}
