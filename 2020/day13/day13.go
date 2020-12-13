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
	scanner := bufio.NewScanner(utils.OpenFile("2020/day13/testinput"))
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
	multiplier := make([]int, 0)
	//currentStep := 1
	for i := 0; i < len(buses)-1; i++ {
		bus1, _ := strconv.Atoi(buses[i])
		var bus2 int
		for {
			var err2 error
			bus2, err2 = strconv.Atoi(buses[i+1])
			if err2 != nil {
				i++
			} else {
				break
			}
		}
		for number := bus1; ; number += bus1 {
			if (number+1)%bus2 == 0 {
				multiplier = append(multiplier, number)
				break
			}
		}
	}
	result := 1
	for _, mult := range multiplier {
		result *= mult
	}
	fmt.Println(multiplier)
	fmt.Println(result)
}
