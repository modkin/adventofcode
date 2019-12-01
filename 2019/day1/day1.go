package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var fuelRequirement = 0

	file, err := os.Open("2019/day1/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		fuelRequirement += i/3 - 2
	}
	fmt.Println("Task 1.1: fuel requirement: ", fuelRequirement)

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
