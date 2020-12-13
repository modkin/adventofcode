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
	fmt.Println(min * busID)
	for _, busString := range buses {
		bus, err := strconv.Atoi(busString)
		if err != nil {
			continue
		}
		fmt.Println(bus)
	}
	start := 1
while:
	for {
		difference := 0
		for _, busString := range buses {
			bus, err := strconv.Atoi(busString)
			if err != nil {
				difference++
				continue
			}
			if (start+difference)%bus != 0 {
				break
			} else {
				difference++
			}
		}
		if difference == len(buses) {
			fmt.Println(start)
			break while
		}
		start++
		if start%1000000000 == 0 {
			fmt.Println(start)
		}
	}
}
