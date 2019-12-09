package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
	"testing"
)

func TestTask1(t *testing.T) {
	test1 := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	test2 := "1102,34915192,34915192,7,4,7,99,0"
	//test3 := "104,1125899906842624,99"

	test1check := []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	test1Split := strings.Split(test1, ",")
	intcode1 := make([]int64, len(test1Split))
	for pos, elem := range test1Split {
		intcode1[pos] = utils.ToInt64(elem)
	}
	test1result := task1(intcode1)
	for i, _ := range intcode1 {
		if test1check[i] != test1result[i] {
			t.Errorf("Wrong!!")
		}
	}

	test2Split := strings.Split(test2, ",")
	intcode2 := make([]int64, len(test2Split))
	for pos, elem := range test2Split {
		intcode1[pos] = utils.ToInt64(elem)
	}
	test2result := task1(intcode2)
	if len(fmt.Sprint(test2result)) != 16 {
		fmt.Println("Error")
	}
}
