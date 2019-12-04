package main

import (
	"fmt"
	"math"
)

func isValidPassword(passwd int) (int, bool) {
	var numbers [6]int
	for i := 0; i < 6; i++ {
		numbers[i] = passwd / int(math.Pow(10, float64(5-i))) % 10
	}
	hasDouble := false
	isIncreasing := true

	pos := 0
	for pos < 5 {
		if numbers[pos] == numbers[pos+1] {
			hasDouble = true
		}
		if numbers[pos] > numbers[pos+1] {
			isIncreasing = false
		}
		pos++
	}
	return passwd, hasDouble && isIncreasing
}

func isValidPassword2(passwd int) (int, bool) {
	var numbers [6]int
	for i := 0; i < 6; i++ {
		numbers[i] = passwd / int(math.Pow(10, float64(5-i))) % 10
	}
	hasDouble := false
	isIncreasing := true

	pos := 0
	for pos < 5 {
		if numbers[pos] == numbers[pos+1] {
			if pos == 0 {
				if numbers[pos+2] != numbers[pos] {
					hasDouble = true
				}
			} else if numbers[pos-1] != numbers[pos] {
				if pos == 4 {
					hasDouble = true
				} else if numbers[pos+1] != numbers[pos+2] {
					hasDouble = true
				}
			}
		}
		if numbers[pos] > numbers[pos+1] {
			isIncreasing = false
		}
		pos++
	}
	return passwd, hasDouble && isIncreasing
}

func main() {
	start := 171309
	end := 643603
	counter := 0
	for i := start; i < end+1; i++ {
		_, correct := isValidPassword(i)
		if correct {
			counter++
		}
	}
	fmt.Println("Task 4.1:", counter)

	counter = 0
	for i := start; i < end; i++ {
		number, correct := isValidPassword2(i)
		if correct {
			fmt.Println(number)
			counter++
		}
	}
	fmt.Println("Task 4.1:", counter)
}
