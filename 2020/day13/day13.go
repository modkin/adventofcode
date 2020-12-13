package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day13/input"))
	scanner.Scan()
	earliest := utils.ToInt(scanner.Text())
	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")
	min, busID := math.MaxInt64, 0
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		wait := utils.ToInt(bus) - (earliest % utils.ToInt(bus))
		if wait < min {
			min = wait
			busID = utils.ToInt(bus)
		}
	}
	fmt.Println("Task 13.1:", min*busID)
	targets := make([][2]int, 0)
	for i, busString := range buses {
		bus, err := strconv.Atoi(busString)
		if err == nil {
			targets = append(targets, [2]int{bus, i})
		}
	}
	startTime := 0
	for i := targets[0][0]; ; i += targets[0][0] {
		if (i+targets[1][1])%targets[1][0] == 0 {
			startTime = i
			break
		}
	}
	step := targets[0][0] * targets[1][0]
	pos := 2
	multiplier := 1
outer:
	for {
		for i := pos; i < len(targets); i++ {
			elem := targets[i]
			if (startTime+elem[1])%elem[0] != 0 {
				startTime += step
				multiplier++
				continue outer
			} else {
				step = step * targets[i][0]
				multiplier = 1
				pos++
			}
			if pos == len(targets) {
				break
			}
		}
		break outer
	}
	fmt.Println("Task 13.2:", startTime)
}
