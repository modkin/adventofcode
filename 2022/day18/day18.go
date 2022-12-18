package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2022/day18/input")

	if err != nil {
		panic(err)
	}

	cubes := make(map[[3]int]bool)
	initCubes := make(map[[3]int]bool)
	airInside := make(map[[3]int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		newCube := [3]int{utils.ToInt(line[0]), utils.ToInt(line[1]), utils.ToInt(line[2])}
		cubes[newCube] = true
		initCubes[newCube] = true
	}
	//fmt.Println(cubes)

	counter := 0
	counter2 := 0
	max := 20
	allPoint := make([][3]int, 0)
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			for z := 0; z < 20; z++ {
				allPoint = append(allPoint, [3]int{x, y, z})
			}
		}
	}
	fmt.Println("Num Cubes 1:", len(cubes))
	fmt.Println("len all point: ", len(allPoint))

	checkDir := false
	xf, yf, zf := true, true, true
	outside := 0

	for _, c := range allPoint {
		if _, ok := cubes[c]; ok {
			continue
		}
		xf, yf, zf = true, true, true
		checkDir = false
		for x := 0; x > -max; x-- {
			check := [3]int{c[0] + x, c[1], c[2]}
			if _, ok := cubes[check]; ok {
				checkDir = true
			}
		}
		if !checkDir {
			xf = false
		}
		checkDir = false
		for x := 0; x < max; x++ {
			check := [3]int{c[0] + x, c[1], c[2]}
			if _, ok := cubes[check]; ok {
				checkDir = true
			}
		}
		if !checkDir {
			xf = false
		}
		checkDir = false
		for y := 0; y > -max; y-- {
			check := [3]int{c[0] + y, c[1], c[2]}
			if _, ok := cubes[check]; ok {
				checkDir = true
			}
		}
		if !checkDir {
			yf = false
		}
		checkDir = false
		for y := 0; y < max; y++ {
			check := [3]int{c[0] + y, c[1], c[2]}
			if _, ok := cubes[check]; ok {
				checkDir = true
			}
		}
		if !checkDir {
			yf = false
		}
		checkDir = false
		for z := 0; z > -max; z-- {
			check := [3]int{c[0] + z, c[1], c[2]}
			if _, ok := cubes[check]; ok {
				checkDir = true
			}
		}
		if !checkDir {
			zf = false
		}
		checkDir = false
		for z := 0; z < max; z++ {
			check := [3]int{c[0] + z, c[1], c[2]}
			if _, ok := cubes[check]; ok {
				checkDir = true
			}
		}
		if !checkDir {
			zf = false
		}
		if xf && yf && zf {
			airInside[c] = true
			//cubes[c] = true
			counter2++
			if _, ok := initCubes[c]; ok {
				fmt.Println("ERROR")
			}
		} else {
			outside++
		}
	}

	fmt.Println("Num cubes 2:", len(cubes))
	fmt.Println("Num air:", len(airInside))
	fmt.Println("outside:", outside)
	fmt.Println("init cubes:", len(initCubes))
	fmt.Println("all:", len(allPoint))
	fmt.Println(len(allPoint) - len(airInside) - outside - len(initCubes))

	for c := range cubes {
		for _, x := range []int{-1, 1} {
			check := [3]int{c[0] + x, c[1], c[2]}
			if _, ok := cubes[check]; !ok {
				if _, ok2 := airInside[check]; !ok2 {
					counter++
				}
			}
		}
		for _, y := range []int{-1, 1} {
			check := [3]int{c[0], c[1] + y, c[2]}
			if _, ok := cubes[check]; !ok {
				if _, ok2 := airInside[check]; !ok2 {
					counter++
				}
			}
		}
		for _, z := range []int{-1, 1} {
			check := [3]int{c[0], c[1], c[2] + z}
			if _, ok := cubes[check]; !ok {
				if _, ok2 := airInside[check]; !ok2 {
					counter++
				}
			}

		}

	}

	fmt.Println(counter2)
	fmt.Println(counter)
	//fmt.Println(counter - counter2*6)
}
