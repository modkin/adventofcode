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
	fmt.Println(earliest)
	fmt.Println(buses)
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
	first := 0
	for i := targets[0][0]; ; i += targets[0][0] {
		if (i+targets[1][1])%targets[1][0] == 0 {
			first = i
			break
		}
	}
	step := targets[0][0] * targets[1][0]
	fmt.Println("First:", first)
	//foundStart := make([]int,len(targets)-1)

outer:
	for {
		for _, elem := range targets {
			if (first+elem[1])%elem[0] != 0 {
				first += step
				//fmt.Println(first)
				continue outer
			}
		}
		fmt.Println("Done:", first)
		break outer

	}

	fmt.Println(targets)

}
