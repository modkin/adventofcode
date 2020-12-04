package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("2020/day4/input")
	if err != nil {
		panic(err)
	}
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	valid := 0
	valid2 := 0
	scanner := bufio.NewScanner(file)
	checkForAll := func(fields []string) {
		all := true
		for _, r := range required {
			if !utils.SliceContains(fields, r) {
				all = false
			}
		}
		if all {
			valid++
		}

	}
	var fields []string
	var fields2 []string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "" {
			checkForAll(fields)
			fields = []string{}

			all2 := true
			for _, r := range required {
				if !utils.SliceContains(fields2, r) {
					all2 = false
				}
			}
			if all2 {
				valid2++
			}
			fields2 = []string{}

		} else {
			for _, elem := range line {
				field := strings.Split(elem, ":")[0]
				fields = append(fields, field)
				data := strings.Split(elem, ":")[1]
				if field == "byr" {
					dataInt, _ := strconv.Atoi(data)
					if dataInt < 1920 || dataInt > 2002 {
						continue
					}
				}
				if field == "iyr" {
					dataInt, _ := strconv.Atoi(data)
					if dataInt < 2010 || dataInt > 2020 {
						continue
					}
				}
				if field == "eyr" {
					dataInt, _ := strconv.Atoi(data)
					if dataInt < 2020 || dataInt > 2030 {
						continue
					}
				}
				if field == "hgt" {
					dataInt, _ := strconv.Atoi(data[0 : len(data)-2])
					if string(data[len(data)-1]) == "m" {
						if dataInt < 150 || dataInt > 193 {
							continue
						}
					} else {
						if dataInt < 59 || dataInt > 76 {
							continue
						}
					}
				}
				if field == "hcl" {
					if string(data[0]) != "#" {
						continue
					}
					if len(data[1:]) != 6 {
						continue
						//TODO: missing
					}
				}
				if field == "ecl" {
					possible := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
					check := false
					for _, p := range possible {
						if data == p {
							check = true
						}
					}
					if !check {
						continue
					}
				}
				if field == "pid" {
					if len(data) != 9 {
						continue
					}
				}
				fields2 = append(fields2, field)
			}
		}

	}
	fmt.Println(valid)
	fmt.Println(valid2)
}
