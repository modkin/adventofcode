package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type oldbrick struct {
	start [3]int
	end   [3]int
	name  string
}

type brick struct {
	parts [][3]int
	name  string
}

type byZ []brick

func (c byZ) Len() int           { return len(c) }
func (c byZ) Less(i, j int) bool { return c[i].parts[0][2] < c[j].parts[0][2] }
func (c byZ) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func touch(first brick, second brick) bool {
	for _, f := range first.parts {
		for _, s := range second.parts {
			if f == s {
				return true
			}
		}
	}
	return false
}

func moveOneDown(b brick) {
	for i, part := range b.parts {
		part[2]--
		b.parts[i] = part
	}
}

func dropBrick(b brick, allB []brick) (brick, bool) {
	droped := false
outer:
	for {
		bNew := brick{slices.Clone(b.parts), b.name}
		moveOneDown(bNew)
		if bNew.parts[0][2] == 0 {
			break
		}
		for _, b2 := range allB {
			if b2.name == b.name {
				continue
			}
			if touch(bNew, b2) {
				break outer
			}
		}
		b = bNew
		droped = true
	}
	return b, droped
}

func main() {
	file, err := os.Open("2023/day22/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var bricks []brick
	name := 65
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "~")
		splitStart := strings.Split(split[0], ",")
		splitEnd := strings.Split(split[1], ",")

		start := [3]int{utils.ToInt(splitStart[0]), utils.ToInt(splitStart[1]), utils.ToInt(splitStart[2])}
		end := [3]int{utils.ToInt(splitEnd[0]), utils.ToInt(splitEnd[1]), utils.ToInt(splitEnd[2])}
		stringName := string(name)

		startX, endX := start[0], end[0]
		startY, endY := start[1], end[1]
		startZ, endZ := start[2], end[2]
		if startX > endX {
			startX, endX = endX, startX
		}
		if startY > endY {
			startY, endY = endY, startY
		}
		if startZ > endZ {
			startZ, endZ = endZ, startZ
		}
		var newBrickParts [][3]int
		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				for z := startZ; z <= endZ; z++ {
					newBrickParts = append(newBrickParts, [3]int{x, y, z})
				}
			}
		}
		bricks = append(bricks, brick{parts: newBrickParts, name: stringName})
		name++
	}
	sort.Sort(byZ(bricks))

	for i, b := range bricks {
		bricks[i], _ = dropBrick(b, bricks)
		sort.Sort(byZ(bricks))
	}

	for _, b := range bricks {
		fmt.Println(b.name)
		fmt.Println(b.parts)
		fmt.Println()
	}

	removeCounter := 0
	removeCount := make(map[string]int)
	var droped bool
	for i, removed := range bricks {
		brickCopy := slices.Clone(bricks)
		oneBrickRemoved := append(brickCopy[:i], brickCopy[(i+1):]...)
		sort.Sort(byZ(bricks))
		for j, b := range oneBrickRemoved {
			oneBrickRemoved[j], droped = dropBrick(b, oneBrickRemoved)
			sort.Sort(byZ(bricks))
			if droped {
				removeCount[removed.name]++
			}
		}
		if removeCount[removed.name] == 0 {
			removeCounter++
		}
	}
	fmt.Println("Day 22.1:", removeCounter)

	sum := 0
	for _, i := range removeCount {
		sum += i
	}
	fmt.Println("Day 22.2:", sum)
}
