package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func checkUpdate(update []int, rules [][2]int) (bool, int) {
	middleElement := update[(len(update)-1)/2]
	for i, curUpdate := range update {
		for _, rule := range rules {
			if curUpdate == rule[0] || curUpdate == rule[1] {
				idx := slices.Index(update, rule[1]) //getIdx(rule[1], update)
				if idx != -1 {
					if idx < i {
						return false, 0
					}
				}
			}
		}
	}
	return true, middleElement
}

func main() {
	file, err := os.Open("2024/day5/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanRules := true
	var orderRules [][2]int
	var pages [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if scanRules {
			if line == "" {
				scanRules = false
				continue
			}
			split := strings.Split(line, "|")
			orderRules = append(orderRules, [2]int{utils.ToInt(split[0]), utils.ToInt(split[1])})
		} else {
			split := strings.Split(line, ",")
			tmp := make([]int, len(split))
			for i, i2 := range split {
				tmp[i] = utils.ToInt(i2)
			}
			pages = append(pages, tmp)
		}
	}
	sum := 0
	var incorrectUpdates [][]int
	for _, page := range pages {
		correct, middle := checkUpdate(page, orderRules)
		//fmt.Println(middle)
		sum += middle
		if !correct {
			incorrectUpdates = append(incorrectUpdates, utils.CopyIntSlice(page))
		}
	}
	fmt.Println("Day 5.1:", sum)
	//fmt.Println(incorrectUpdates)
	sum2 := 0
	cmp := func(a, b int) int {
		for _, rule := range orderRules {
			if rule[0] == a && rule[1] == b {
				return -1
			} else if rule[1] == a && rule[0] == b {
				return 1
			}
		}
		return 0
	}
	for _, update := range incorrectUpdates {
		slices.SortFunc(update, cmp)
		sum2 += update[(len(update)-1)/2]
	}

	fmt.Println("Day 5.2:", sum2)
}
