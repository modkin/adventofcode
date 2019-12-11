package main

import "testing"

func TestRotation(t *testing.T) {
	direction := [2]int{0, 1}
	direction = rotate90left(direction)
	if direction != [2]int{-1, 0} {
		panic("Error")
	}
	direction = rotate90left(direction)
	if direction != [2]int{0, -1} {
		panic("Error")
	}
	direction = rotate90left(direction)
	if direction != [2]int{1, 0} {
		panic("Error")
	}
	direction = rotate90left(direction)
	if direction != [2]int{0, 1} {
		panic("Error")
	}

	direction = rotate90right(direction)
	if direction != [2]int{1, 0} {
		panic("Error")
	}
	direction = rotate90right(direction)
	if direction != [2]int{0, -1} {
		panic("Error")
	}
	direction = rotate90right(direction)
	if direction != [2]int{-1, 0} {
		panic("Error")
	}
	direction = rotate90right(direction)
	if direction != [2]int{0, 1} {
		panic("Error")
	}

}
