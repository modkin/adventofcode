package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printDisks(in []int) {
	for _, n := range in {
		if n == -1 {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", n)
		}
	}
}

func countMinOne(in []int) int {
	count := 0
	for _, n := range in {
		if n == -1 {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open("2024/day9/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	id := 0
	isDisk := true
	dist := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		for _, s := range split {
			if isDisk {
				for i := 0; i < utils.ToInt(s); i++ {
					dist = append(dist, id)
				}
			} else {
				for i := 0; i < utils.ToInt(s); i++ {
					dist = append(dist, -1)
				}

			}
			if isDisk {
				id++
			}
			isDisk = !isDisk
		}
	}

	//printDisks(dist)

	for countMinOne(dist) > 0 {
	outer:
		for insertIdx, n := range dist {
			if n == -1 {
				for i := len(dist) - 1; i > 0; i-- {
					if dist[insertIdx] == -1 {
						dist[insertIdx] = dist[i]
						dist = dist[:i]
						break outer
					}
				}

			}
		}
	}
	checksum := 0
	for i, n := range dist {
		checksum += i * n
	}
	fmt.Println(checksum)
}
