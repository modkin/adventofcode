package main

import (
	"testing"
)

func TestTask1(t *testing.T) {
	ret := runTask("./testInput1")
	if ret != 86 {
		t.Errorf("Test 1 wrong %d != 86", ret)
	}
}
func TestTask2(t *testing.T) {
	ret := runTask("./testInput2")
	if ret != 136 {
		t.Errorf("Test 1 wrong %d != 136", ret)
	}
}
func TestTask3(t *testing.T) {
	ret := runTask("./testInput3")
	if ret != 81 {
		t.Errorf("Test 1 wrong %d != 81", ret)
	}
}
