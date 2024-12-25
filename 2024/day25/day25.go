package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.ReadFileIntoLines("2024/day25/input")

	schemes := []map[[2]int]string{}
	tmp := make(map[[2]int]string)
	y := 0
	ymax := 7
	xmax := 6
	for _, line := range lines {
		if line == "" {
			schemes = append(schemes, tmp)
			tmp = make(map[[2]int]string)
			y = 0
		} else {
			for x, s := range strings.Split(line, "") {
				tmp[[2]int{x, y}] = s

			}
			y++
		}

	}
	schemes = append(schemes, tmp)

	getLine := func(y int, schceme map[[2]int]string) string {
		var out string
		for x := 0; x < 5; x++ {
			out += schceme[[2]int{x, y}]
		}
		return out
	}

	locks := []map[[2]int]string{}
	keys := []map[[2]int]string{}

	for _, scheme := range schemes {
		line := getLine(0, scheme)
		fmt.Println(line)
		if strings.Count(line, "#") == xmax-1 {
			locks = append(locks, scheme)
		} else {
			keys = append(keys, scheme)
		}
	}

	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			fits := true
			for x := 0; x < xmax; x++ {
				for y = 0; y < ymax; y++ {
					if lock[[2]int{x, y}] == "#" && key[[2]int{x, y}] == "#" {
						fits = false
					}
				}
			}
			if fits {
				count++
			}
		}

	}
	fmt.Println(count)
}
