package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"regexp"
)

type machine struct {
	a      [2]int
	b      [2]int
	price  [2]int
	price2 [2]int
}

func solve(ma machine, part1 bool) (bool, [2]int) {
	var p [2]int
	if part1 {
		p = ma.price
	} else {
		p = ma.price2
	}
	det := ma.a[0]*ma.b[1] - ma.a[1]*ma.b[0]
	tmp1 := p[0]*ma.b[1] - ma.b[0]*p[1]
	tmp2 := ma.a[0]*p[1] - p[0]*ma.a[1]
	aPush := tmp1 / det
	bPush := tmp2 / det
	if float64(aPush) == float64(tmp1)/float64(det) && float64(bPush) == float64(tmp2)/float64(det) {
		return true, [2]int{aPush, bPush}
	} else {
		return false, [2]int{aPush, bPush}
	}
}

func main() {
	file, err := os.ReadFile("2024/day13/input")
	if err != nil {
		panic(err)
	}

	fullReg := regexp.MustCompile(`.*\+(\d*), Y.?(\d*)\n.*\+(\d*), Y.?(\d*)\n.*=(\d*), Y=(\d*)`)

	var allMachines []machine

	for _, strings := range fullReg.FindAllStringSubmatch(string(file), -1) {
		a := [2]int{utils.ToInt(strings[1]), utils.ToInt(strings[2])}
		b := [2]int{utils.ToInt(strings[3]), utils.ToInt(strings[4])}
		price := [2]int{utils.ToInt(strings[5]), utils.ToInt(strings[6])}
		price2 := [2]int{utils.ToInt(strings[5]) + 10000000000000, utils.ToInt(strings[6]) + 10000000000000}
		newMach := machine{a: a, b: b, price: price, price2: price2}
		allMachines = append(allMachines, newMach)
	}

	sum1 := 0
	sum2 := 0

	for _, ma := range allMachines {
		intSol, pushes := solve(ma, true)
		if intSol {
			sum1 += pushes[0]*3 + pushes[1]
		}

		intSol, pushes = solve(ma, false)
		if intSol {
			sum2 += pushes[0]*3 + pushes[1]
		}
	}
	fmt.Println("Day 13.1:", sum1)
	fmt.Println("Day 13.2:", sum2)
}
