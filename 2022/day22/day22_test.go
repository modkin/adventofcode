package main

import "testing"

func TestRotation(t *testing.T) {
	direction := [2]int{0, 1}
	direction = turnLeft(direction)
	if direction != [2]int{1, 0} {
		panic("Error")
	}
	direction = turnLeft(direction)
	if direction != [2]int{0, -1} {
		panic("Error")
	}
	direction = turnLeft(direction)
	if direction != [2]int{-1, 0} {
		panic("Error")
	}
	direction = turnLeft(direction)
	if direction != [2]int{0, 1} {
		panic("Error")
	}

	direction = turnRight(direction)
	if direction != [2]int{-1, 0} {
		panic("Error")
	}
	direction = turnRight(direction)
	if direction != [2]int{0, -1} {
		panic("Error")
	}
	direction = turnRight(direction)
	if direction != [2]int{1, 0} {
		panic("Error")
	}
	direction = turnRight(direction)
	if direction != [2]int{0, 1} {
		panic("Error")
	}

}
