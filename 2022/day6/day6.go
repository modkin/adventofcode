package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(array []rune, toCount rune) int {
	counter := 0
	for _, r := range array {
		if r == toCount {
			counter++
		}
	}
	return counter
}

func check(in []rune) bool {
	for _, i2 := range in {
		if count(in, i2) > 1 {
			return false
		}
	}
	return true
}

func main() {

	file, err := os.Open("2022/day6/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	signal := make([]rune, 0)

	for scanner.Scan() {
		signal = []rune(scanner.Text())

	}

	fmt.Print("Day 6.1: ")
	for i := range signal {
		if check(signal[i : i+4]) {
			fmt.Println(i + 4)
			break
		}
	}

	fmt.Print("Day 6.2: ")
	for i := range signal {
		if check(signal[i : i+14]) {
			fmt.Println(i + 14)
			break
		}
	}

}
