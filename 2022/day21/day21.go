package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	number    int
	input     [2]string
	operation string
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

	for scanner.Scan() {
		input := [2]string{}
		operation := ""
		number := 0
		split := strings.Split(scanner.Text(), " ")
		name := strings.Trim(split[0], ":")

		if ret, err2 := strconv.Atoi(split[1]); err2 != nil {
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

	for {
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
			break
		}
	}
}
