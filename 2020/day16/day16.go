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

func findValidCategories(numbers []int, validFields map[string]map[int]bool) map[string]bool {
	ret := make(map[string]bool)
outer:
	for c, cat := range validFields {
		for _, elem := range numbers {
			if !cat[elem] {
				continue outer
			}
		}
		ret[c] = true
	}
	return ret
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
	var numberOfCategories int
	ownTicket := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "nearby tickets:" {
			break
		}
		if line == "" {
			continue
		}
		ownTicket = strings.Split(line, ",")
		numberOfCategories = len(ownTicket)
	}

	countInvalid := 0
	numbersInCategory := make([][]int, numberOfCategories)
scan:
	for scanner.Scan() {
		fieldSlice := strings.Split(scanner.Text(), ",")
		for _, elem := range fieldSlice {
			if !checkValid(utils.ToInt(elem), fields) {
				countInvalid += utils.ToInt(elem)
				continue scan
			}
		}

		for i, f := range fieldSlice {
			numbersInCategory[i] = append(numbersInCategory[i], utils.ToInt(f))
		}
	}
	fmt.Println("Task 16.1:", countInvalid)
	dep := 1
	categoryToPos := make(map[string]int)
	allPossibleCat := make([]map[string]bool, 0)
	for _, elem := range numbersInCategory {
		tmp := findValidCategories(elem, fields)
		allPossibleCat = append(allPossibleCat, tmp)
	}
	for i := 0; i < len(ownTicket); i++ {
		var currentCat string
		for i, elem := range allPossibleCat {
			if len(elem) == 1 {
				for k := range elem {
					// only one iteration
					currentCat = k
					categoryToPos[k] = i
				}
			}
		}
		for _, elem := range allPossibleCat {
			delete(elem, currentCat)
		}
	}
	for cat, pos := range categoryToPos {
		if strings.Split(cat, " ")[0] == "departure" {
			dep *= utils.ToInt(ownTicket[pos])
		}
	}
	fmt.Println("Task 16.2:", dep)
}
