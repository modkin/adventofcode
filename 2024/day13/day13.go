package main

import (
	"adventofcode/utils"
	"fmt"
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
	lines := utils.ReadFileIntoLines("2024/day13/input")

	mach := regexp.MustCompile(`Button .: X(.\d*), Y(.\d*).*`)
	regPrice := regexp.MustCompile(`.*X.(\d*), Y.(\d*).*`)

	var allMachines []machine
	for i := 0; i < len(lines); i += 4 {

		coordsA := mach.FindStringSubmatch(lines[i+0])
		coordsB := mach.FindStringSubmatch(lines[i+1])
		a := [2]int{utils.ToInt(coordsA[1]), utils.ToInt(coordsA[2])}
		b := [2]int{utils.ToInt(coordsB[1]), utils.ToInt(coordsB[2])}
		coordsPrice := regPrice.FindStringSubmatch(lines[i+2])
		price := [2]int{utils.ToInt(coordsPrice[1]), utils.ToInt(coordsPrice[2])}
		price2 := [2]int{utils.ToInt(coordsPrice[1]) + 10000000000000, utils.ToInt(coordsPrice[2]) + 10000000000000}
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
