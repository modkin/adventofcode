package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
)

type machine struct {
	a     [2]int
	b     [2]int
	price [2]int
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
		newMach := machine{a: a, b: b, price: price}
		allMachines = append(allMachines, newMach)
	}

	sum := 0

	//ma := allMachines[0]
	//xCoord := ma.a[0]*80 + ma.b[0]*40
	//yCoord := ma.a[1]*80 + ma.b[1]*40
	//fmt.Println(ma.price == [2]int{xCoord, yCoord})

	for _, ma := range allMachines {
		minPush := math.MaxInt
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				xCoord := ma.a[0]*a + ma.b[0]*b
				yCoord := ma.a[1]*a + ma.b[1]*b
				if ma.price == [2]int{xCoord, yCoord} {
					if cost := a*3 + b; cost < minPush {
						minPush = cost
						fmt.Println(a, b)
					}
				}
			}
		}
		if minPush != math.MaxInt {
			sum += minPush
		}
	}
	fmt.Println(sum)
}
