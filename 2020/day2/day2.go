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
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		min, _ := strconv.Atoi(strings.Split(line[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(line[0], "-")[1])
		require := strings.Split(line[1], "")[0]
		password := line[2]
		amount := strings.Count(password, require)
		//fmt.Println(min, max, require, password, amount)
		if amount >= min && amount <= max {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
}
