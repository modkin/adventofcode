package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type multipleBags struct {
	color  string
	amount int
}

func countInsideBags(bag multipleBags, possible map[string][]multipleBags, total int) int {
	for _, value := range possible[bag.color] {
		tmp := countInsideBags(value, possible, 0)
		total += value.amount
		if tmp != 0 {
			total += tmp * value.amount
		}
	}
	return total
}

func findShiny(bags []multipleBags, possible map[string][]multipleBags) bool {
	tmp := false
	for _, elem := range bags {
		if elem.color == "shiny gold" {
			tmp = true
		} else {
			tmp = tmp || findShiny(possible[elem.color], possible)
		}
	}
	return tmp
}

func main() {
	file, err := os.Open("2020/day7/input")
	if err != nil {
		panic(err)
	}

	possbile := make(map[string][]multipleBags)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " bags contain ")
		bag := line[0]
		contents := strings.Split(line[1], ",")
		inside := make([]multipleBags, 0)
		for _, c := range contents {
			tmp := strings.Split(strings.TrimSpace(c), " ")
			if tmp[1] != "other" {
				amount, err := strconv.Atoi(tmp[0])
				if err != nil {
					panic(err)
				}
				inside = append(inside, multipleBags{tmp[1] + " " + tmp[2], amount})
			}
		}
		possbile[bag] = inside
	}

	totalRed := 0
	for _, content := range possbile {
		if findShiny(content, possbile) {
			totalRed++
		}
	}
	fmt.Println("Task 7.1:", totalRed)
	fmt.Println("Task 7.2:", countInsideBags(multipleBags{"shiny gold", 0}, possbile, 0))
}
