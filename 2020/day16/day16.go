package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func checkValid(field int, validFields map[string]map[int]bool) bool {
	for _, elem := range validFields {
		if elem[field] {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day16/input"))

	fields := make(map[string]map[int]bool)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "your ticket:" {
			break
		}
		if line == "" {
			continue
		}
		rule := strings.Split(line, ":")[0]
		fields[rule] = make(map[int]bool)
		ranges := strings.Split(strings.Split(line, ":")[1], "or")
		for _, elem := range ranges {
			tmp := strings.Split(strings.TrimSpace(elem), "-")
			for i := utils.ToInt(tmp[0]); i <= utils.ToInt(tmp[1]); i++ {
				fields[rule][i] = true
			}
		}

	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "nearby tickets:" {
			break
		}
		if line == "" {
			continue
		}
	}

	countInvalid := 0

	for scanner.Scan() {
		fieldSlice := strings.Split(scanner.Text(), ",")
		for _, elem := range fieldSlice {
			if !checkValid(utils.ToInt(elem), fields) {
				countInvalid += utils.ToInt(elem)
			}
		}
	}
	fmt.Println(countInvalid)
}
