package main

import "testing"

func TestTask1(t *testing.T) {
	task1result := task1()
	if task1result != 18812 {
		t.Errorf("Task 1 wrong %d != 18812", task1result)
	}

	task2result := task2()
	correct := 25534964
	if task2result != correct {
		t.Errorf("Task 1 wrong %d != %d", task2result, correct)
	}
}
