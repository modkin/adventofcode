package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("2020/day2/input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	validPasswords := 0
	validPasswords2 := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		min, _ := strconv.Atoi(strings.Split(line[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(line[0], "-")[1])
		require := strings.Split(line[1], "")[0]
		password := line[2]
		amount := strings.Count(password, require)
		if amount >= min && amount <= max {
			validPasswords++
		}
		passwordSplit := strings.Split(password, "")
		if (passwordSplit[min-1] == require) != (passwordSplit[max-1] == require) {
			validPasswords2++
		}
	}
	fmt.Println("Task 2.1: ", validPasswords)
	fmt.Println("Task 2.2: ", validPasswords2)
}
