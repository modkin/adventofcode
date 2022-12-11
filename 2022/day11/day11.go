package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type CPU struct {
	x int
}

type monkey struct {
	items     []int
	operation func(int) int
	testWorry func(int) int
}

func main() {

	file, err := os.Open("2022/day11/input")
	if err != nil {
		panic(err)
	}

	//grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)

	monkeys := make([]monkey, 0)

	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		if tmp[0] == "Monkey" {
			scanner.Scan()
			tmp = strings.Split(strings.TrimSpace(scanner.Text()), " ")
			startItems := make([]int, 0)
			for _, item := range tmp[2:] {
				startItems = append(startItems, utils.ToInt(strings.ReplaceAll(item, ",", "")))
			}
			scanner.Scan()
			tmp = strings.Split(strings.TrimSpace(scanner.Text()), " ")
			op := tmp[4]
			op2 := tmp[5]
			operation := func(old int) int {

				if op == "+" {
					if op2 == "old" {
						return old + old
					} else {
						return old + utils.ToInt(op2)
					}
				} else {
					if op2 == "old" {
						return old * old
					} else {
						return old * utils.ToInt(op2)
					}
				}
			}

			scanner.Scan()
			tmp = strings.Split(strings.TrimSpace(scanner.Text()), " ")
			divby := utils.ToInt(tmp[3])
			scanner.Scan()
			tmp = strings.Split(strings.TrimSpace(scanner.Text()), " ")
			ifTrue := utils.ToInt(tmp[5])
			scanner.Scan()
			tmp = strings.Split(strings.TrimSpace(scanner.Text()), " ")
			ifFalse := utils.ToInt(tmp[5])
			testWorry := func(worryLevel int) int {
				if worryLevel%divby == 0 {
					return ifTrue
				} else {
					return ifFalse
				}
			}
			newMonkey := monkey{startItems, operation, testWorry}
			monkeys = append(monkeys, newMonkey)
		}
	}
	scores := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for mi, m := range monkeys {
			for _, item := range m.items {
				scores[mi]++
				item = m.operation(item)
				item = item / 3
				targetMonkey := m.testWorry(item)
				monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, item)
				if len(monkeys[mi].items) == 1 {
					monkeys[mi].items = make([]int, 0)
				} else {
					monkeys[mi].items = monkeys[mi].items[1:]
				}
			}
		}
		fmt.Println("Round", round)
		for i, m := range monkeys {
			fmt.Println(i, m.items)
		}
	}
	fmt.Println(scores)

	sort.Ints(scores)
	monkeyBuisness := scores[len(scores)-1] * scores[len(scores)-2]
	fmt.Println("Day 11.1:", monkeyBuisness)

}
