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
	startTime, pos := 0, 1
	step := targets[0][0]
outer:
	for pos != len(targets) {
		for i := pos; i < len(targets); i++ {
			elem := targets[i]
			if (startTime+elem[1])%elem[0] != 0 {
				startTime += step
				continue outer
			} else {
				step = step * targets[i][0]
				pos++
			}
		}
	}
	fmt.Println("Task 13.2:", startTime)
}
