package main

import "testing"

func TestTask1(t *testing.T) {
	task1result := task1()
	if task1result != 7988899 {
		t.Errorf("Task 1 wrong %d != 7988899", task1result)
	}
}
