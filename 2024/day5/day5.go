package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getIdx(in int, update []int) int {
	for i, i2 := range update {
		if i2 == in {
			return i
		}
	}
	return -1
}

func checkUpdate(update []int, rules [][2]int) int {
	middleElement := update[(len(update)-1)/2]
	for i, curUpdate := range update {
		for _, rule := range rules {
			if curUpdate == rule[0] || curUpdate == rule[1] {
				idx := getIdx(rule[1], update)
				if idx != -1 {
					if idx < i {
						return 0
					}
				}
			}
		}
	}
	return middleElement
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
	for _, page := range pages {
		sum += checkUpdate(page, orderRules)
	}

	fmt.Println(orderRules)
	fmt.Println(pages)
	fmt.Println(sum)
}
