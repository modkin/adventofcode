package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(left [2]int, right [2]int) [2]int {
	return [2]int{left[0] + right[0], left[1] + right[1]}
}

func hash(input map[[2]int]string, maxX int, maxY int) string {
	out := ""
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			out += input[[2]int{x, y}]
		}
	}
	return out
}

func unpack(input string, maxX int, maxY int) map[[2]int]string {
	ret := make(map[[2]int]string)
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			ret[[2]int{x, y}] = string(input[y*(maxX+1)+x])
		}
	}
	return ret
}

func main() {
	file, err := os.Open("2023/day14/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	//var lines []string
	//for scanner.Scan() {
	//	lines = append(lines, scanner.Text())
	//
	//}

	cache := make(map[string]string)
	table := make(map[[2]int]string)
	y := 0
	var maxX int
	var maxY int
	for scanner.Scan() {
		for x, pipe := range scanner.Text() {
			table[[2]int{x, y}] = string(pipe)
		}
		y++
		maxX = len(scanner.Text()) - 1
	}
	maxY = y - 1
	fmt.Println(maxX, maxY)

	tilt := func(tab string, move [2]int) string {
		hashed := tab + strconv.Itoa(move[0]) + strconv.Itoa(move[1])
		if val, ok := cache[hashed]; ok {
			return val[:(maxX+1)*(maxY+1)]
		} else {
			newTable := unpack(tab, maxX, maxY)
			for {
				counter := 0
				for pos, s := range newTable {
					if s == "O" {
						newPos := sum(pos, move)
						if newTable[newPos] == "." {
							newTable[newPos] = "O"
							newTable[pos] = "."
							counter++
						}
					}
				}
				if counter == 0 {
					break
				}
			}
			hashedTable := hash(newTable, maxX, maxY)

			newHashed := hashedTable + strconv.Itoa(move[0]) + strconv.Itoa(move[1])
			cache[hashed] = newHashed
			return hashedTable
		}
	}
	utils.Print2DStringsGrid(table)
	hashedTable := hash(table, maxX, maxY)
	hashedTable = tilt(hashedTable, [2]int{0, -1})
	getLoad := func() int {
		load := 0
		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
				if table[[2]int{x, y}] == "O" {
					load += (maxY + 1) - y
				}
			}
		}
		return load
	}
	table = unpack(hashedTable, maxX, maxY)
	fmt.Println(getLoad())

	cycles := 1000000000
	sample := ""
	for i := 0; i < cycles; i++ {
		if i != 0 {
			hashedTable = tilt(hashedTable, [2]int{0, -1})
		}
		hashedTable = tilt(hashedTable, [2]int{-1, 0})
		hashedTable = tilt(hashedTable, [2]int{0, 1})
		hashedTable = tilt(hashedTable, [2]int{1, 0})
		if i%1000000 == 0 {
			if i > 0 && sample == "" {
				sample = hashedTable
				i += 7 * 142000000
			}
			fmt.Println(i)
		}

	}
	table = unpack(hashedTable, maxX, maxY)
	//utils.Print2DStringsGrid(table)
	fmt.Println(getLoad())

}
