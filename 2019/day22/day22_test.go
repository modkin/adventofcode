package main

import (
	"testing"
)

func TestMain(m *testing.M) {
	main()
}

func TestTask1(t *testing.T) {
	positions := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	revPositions := []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	length := int64(10)
	for i, pos := range positions {
		tmp := trackReverse(pos, length)
		if tmp != revPositions[i] {
			t.Error("Wrong")
		}
	}
	cut3Positions := []int64{7, 8, 9, 0, 1, 2, 3, 4, 5, 6}

	for i, pos := range positions {
		tmp := trackCut(3, pos, length)
		if tmp != cut3Positions[i] {
			t.Error("Wrong", pos)
		}
	}

	cut3negPositions := []int64{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}

	for i, pos := range positions {
		tmp := trackCut(-4, pos, length)
		if tmp != cut3negPositions[i] {
			t.Error("Wrong", pos)
		}
	}

	dealWithInc3 := []int64{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}

	for i, pos := range positions {
		tmp := trackdealWithInc(3, pos, length)
		if tmp != dealWithInc3[i] {
			t.Error("Wrong", pos)
		}
	}
}
