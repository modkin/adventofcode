package main

import (
	"adventofcode/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("2015/day6/input")
	if err != nil {
		panic(err)
	}

	wires := make(map[string]func() uint8)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		inst := strings.Split(line, " ")
		if strings.Contains(line, "RSHIFT") {
			tmp := func() uint8 {
				return wires[inst[0]]() >> uint(utils.ToInt(inst[2]))
			}
			wires[inst[4]] = tmp
		} else if strings.Contains(line, "LSHIFT") {

		} else if strings.Contains(line, "OR") {

		} else if strings.Contains(line, "AND") {

		} else if strings.Contains(line, "NOT") {

		}
	}

	source := inst[0]

	tmp := func() uint8 {
		return 8
	}
	wires["b"] = tmp
	wires["b"]()
}
