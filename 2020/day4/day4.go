package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("2020/day4/input")
	if err != nil {
		panic(err)
	}
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var fields []string
	valid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "" {
			all := true
			for _, r := range required {
				if !utils.SliceContains(fields, r) {
					all = false
				}
			}
			if all {
				valid++
			}
			fields = []string{}

		} else {
			for _, elem := range line {
				field := strings.Split(elem, ":")[0]
				fields = append(fields, field)
			}
		}

	}
	fmt.Println(valid)
}
