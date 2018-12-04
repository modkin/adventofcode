package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"fmt"
)

func main() {
	fmt.Println("Day1.1: ", day1.SumFile("./day1input.txt"))
	fmt.Println("Day1.2: ", day1.FindFirstTwice("./day1input.txt"))
	fmt.Println("Day2.1: ", day2.Checksum("./day2input.txt"))
}
