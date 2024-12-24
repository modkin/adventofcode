package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type gate struct {
	in  [2]string
	out string
	op  string
}

func performOp(values map[string]int, g gate) bool {
	if _, ok := values[g.in[0]]; !ok {
		return true
	}
	if _, ok := values[g.in[1]]; !ok {
		return true
	}
	if g.op == "AND" {
		values[g.out] = values[g.in[0]] & values[g.in[1]]
	} else if g.op == "OR" {
		values[g.out] = values[g.in[0]] | values[g.in[1]]
	} else if g.op == "XOR" {
		values[g.out] = values[g.in[0]] ^ values[g.in[1]]
	}
	return false
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day24/input")

	values := make(map[string]int)
	gates := []gate{}

	stillValues := true
	maxZ := 0

	for _, line := range lines {
		if line == "" {
			stillValues = false
			continue
		}
		if stillValues {
			split := strings.Split(line, " ")
			values[strings.Trim(split[0], ":")] = utils.ToInt(split[1])

		} else {
			splitLine := strings.Split(line, " ")
			newGate := gate{[2]string{splitLine[0], splitLine[2]}, splitLine[4], splitLine[1]}
			gates = append(gates, newGate)
		}
	}

	fmt.Println(values)
	fmt.Println(gates)

	goOn := true
	for goOn {
		goOn = false
		for _, g := range gates {
			if performOp(values, g) {
				goOn = true
			}
		}
	}
	fmt.Println(values)

	for s, _ := range values {
		if string(s[0]) == "z" {
			maxZ++
		}
	}

	fmt.Println()
	var out string
	for z := maxZ - 1; z >= 0; z-- {
		if z < 10 {
			fmt.Print(values["z0"+strconv.Itoa(z)])
			out += strconv.Itoa(values["z0"+strconv.Itoa(z)])
		} else {
			fmt.Print(values["z"+strconv.Itoa(z)])
			out += strconv.Itoa(values["z"+strconv.Itoa(z)])
		}
	}
	part1, _ := strconv.ParseInt(out, 2, 64)
	fmt.Println()
	fmt.Println(part1)
}
