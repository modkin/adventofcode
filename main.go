package main

import (
	"adventofcode/day1"
	"adventofcode/day10"
	"adventofcode/day2"
	"adventofcode/day3"
	"adventofcode/day4"
	"adventofcode/day5"
	"fmt"
)

func main() {
	fmt.Println("Day1.1:", day1.SumFile("./day1input.txt"))
	fmt.Println("Day1.2:", day1.FindFirstTwice("./day1input.txt"))
	fmt.Println("Day2.1:", day2.Checksum("./day2input.txt"))
	fmt.Println("Day2.2:", day2.FindSpecial("./day2input.txt"))
	fmt.Println("Day3.1:", day3.Task1())
	fmt.Println("Day3.2:", day3.Task2())
	fmt.Println("Day4.1:", day4.Task1())
	fmt.Println("Day4.2:", day4.Task2())
	fmt.Println("Day5.1:", day5.Task1())
	fmt.Println("Day5.2:", day5.Task2())
	day10.Task1()
}
