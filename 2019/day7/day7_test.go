package main

import (
	"adventofcode/utils"
	"io/ioutil"
	"strings"
	"testing"
)

func TestTask1(t *testing.T) {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt(elem)
	}
	task1result := task1(intcode)
	if task1result != 18812 {
		t.Errorf("Task 1 wrong %d != 18812", task1result)
	}

	task2result := task2(intcode)
	correct := 25534964
	if task2result != correct {
		t.Errorf("Task 1 wrong %d != %d", task2result, correct)
	}
}
