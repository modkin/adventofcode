package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type op struct {
	name      string
	operation string
	result    int
	left      *op
	right     *op
}

type monkey struct {
	number    int
	input     [2]string
	operation string
}

func findDep(mon monkey, monkeys map[string]monkey, dep string) bool {
	if mon.input[0] == "" {
		return false
	} else if mon.input[0] == dep || mon.input[1] == dep {
		return true
	} else {
		return findDep(monkeys[mon.input[0]], monkeys, dep) || findDep(monkeys[mon.input[1]], monkeys, dep)
	}
}

func calc(mon monkey, first, second int) int {
	if mon.operation == "+" {
		return first + second
	}
	if mon.operation == "-" {
		return first - second
	}
	if mon.operation == "/" {
		return first / second
	}
	if mon.operation == "*" {
		return first * second
	}
	return math.MaxInt
}

func main() {
	file, err := os.Open("2022/day21/input")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	monkeys := make(map[string]monkey)
	allVar := make(map[string]bool)

	for scanner.Scan() {
		input := [2]string{}
		operation := ""
		number := 0
		split := strings.Split(scanner.Text(), " ")
		name := strings.Trim(split[0], ":")

		if ret, err2 := strconv.Atoi(split[1]); err2 != nil {
			for _, tmp := range []string{split[1], split[3]} {
				if _, ok := allVar[tmp]; ok {
					fmt.Println("DOUBLE", tmp)
				} else {
					allVar[tmp] = true
				}
			}
			input = [2]string{split[1], split[3]}
			operation = split[2]
		} else {
			number = ret
		}
		newMonkey := monkey{
			number, input, operation,
		}
		monkeys[name] = newMonkey
	}

	part1Sol := 0
	counter := 0

	for {
		counter++
		for name, mon := range monkeys {
			if mon.number == 0 {
				first := monkeys[mon.input[0]].number
				second := monkeys[mon.input[1]].number
				if first != 0 && second != 0 {
					mon.number = calc(mon, first, second)
				}
			}
			monkeys[name] = mon
		}
		if monkeys["root"].number != 0 {
			fmt.Println(monkeys["root"].number)
			part1Sol = monkeys["root"].number
			break
		}
	}
	fmt.Println("Day 21.1:", part1Sol)

	root0 := monkeys["root"].input[0]
	root1 := monkeys["root"].input[1]
	var target string
	if findDep(monkeys[root0], monkeys, "humn") {
		mon := monkeys[root0]
		mon.number = monkeys[root1].number
		monkeys[root0] = mon
		fmt.Println("equal:", mon.number)
		target = root0
	}
	if findDep(monkeys[root1], monkeys, "humn") {
		mon := monkeys[root1]
		mon.number = monkeys[root0].number
		monkeys[root1] = mon
		fmt.Println("equal:", mon.number)
		target = root1
	}
	fmt.Println(target)
	cur := target
	targetNum := monkeys[target].number
	for {
		mon := monkeys[cur]

		for i, s := range mon.input {
			if s != "humn" {
				if !findDep(monkeys[s], monkeys, "humn") {
					if mon.operation == "+" {
						targetNum -= monkeys[s].number
					}
					if mon.operation == "-" {
						if i == 1 {
							targetNum += monkeys[s].number
						} else {
							targetNum = monkeys[s].number - targetNum
						}
					}
					if mon.operation == "*" {
						targetNum /= monkeys[s].number
					}
					if mon.operation == "/" {
						if i == 1 {
							targetNum *= monkeys[s].number
						} else {
							targetNum = monkeys[s].number / targetNum
						}
					}
				} else {
					cur = s
				}
			} else {
				cur = "humn"
			}
		}
		if cur == "humn" {
			fmt.Println("Day 21.2:", targetNum)
			break
		}
	}
}
