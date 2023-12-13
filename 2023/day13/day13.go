package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("2023/day13/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	var mirrors [][]string

	fmt.Println(lines)
	newMirror := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			mirrors = append(mirrors, newMirror)
			newMirror = make([]string, 0)
		} else {
			newMirror = append(newMirror, line)
		}
	}
	mirrors = append(mirrors, newMirror)
	fmt.Println(mirrors)

	findMirror := func(mirror []string) (int, bool) {

		for col := 0; col < len(mirror[0])-1; col++ {
			startLeft := col
			startRight := col + 1
			isMirror := true
		colLoop:
			for y := 0; y < len(mirror); y++ {
				for x := 0; x < len(mirror[0]); x++ {
					if startLeft-x < 0 || startRight+x >= len(mirror[0]) {
						continue colLoop
					}
					if mirror[y][startLeft-x] != mirror[y][startRight+x] {
						isMirror = false
					}
				}
			}
			if isMirror {
				fmt.Println("Col Mirror:", col+1)
				return col + 1, true
			}
			startLeft++
			startRight++
		}
		for row := 0; row < len(mirror)-1; row++ {
			startTop := row
			startBot := row + 1
			isMirror := true
		rowLoop:
			for x := 0; x < len(mirror[0]); x++ {
				for y := 0; y < len(mirror); y++ {
					if startTop-y < 0 || startBot+y >= len(mirror) {
						continue rowLoop
					}
					if mirror[startTop-y][x] != mirror[startBot+y][x] {
						isMirror = false
					}
				}
			}
			if isMirror {
				fmt.Println("Row Mirror:", row+1)
				return row + 1, false
			}
			startTop++
			startBot++
		}

		return math.MaxInt, false
	}
	ret := 0
	for _, mirror := range mirrors {
		nbr, col := findMirror(mirror)
		if col {
			ret += nbr
		} else {
			ret += (100 * nbr)
		}
	}
	fmt.Println(ret)
}
