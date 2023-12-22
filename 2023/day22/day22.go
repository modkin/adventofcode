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

func getFallDepth(xStart int, yStart int, zStart int, bricks []oldbrick, ignore oldbrick) int {
	endHeight := 0
	for _, b := range bricks {
		if b == ignore {
			continue
		}
		if zStart < b.start[2] && zStart < b.end[2] {
			continue
		}
		higherZ := b.start[2]
		if b.end[2] > higherZ {
			higherZ = b.end[2]
		}
		var xStep int
		var yStep int
		if b.end[0]-b.start[0] == 0 {
			xStep = 1
		} else {
			xStep = (b.end[0] - b.start[0]) / utils.IntAbs(b.end[0]-b.start[0])
		}
		if b.end[1]-b.start[1] == 0 {
			yStep = 1
		} else {
			yStep = (b.end[1] - b.start[1]) / utils.IntAbs(b.end[1]-b.start[1])
		}
		for x := b.start[0]; x <= b.end[0]; x += xStep {
			for y := b.start[1]; y <= b.end[1]; y += yStep {
				if x == xStart && y == yStart {
					if higherZ > endHeight {
						endHeight = higherZ
					}
				}
			}
		}
	}
	if endHeight == 0 {
		return zStart
	} else {
		return zStart - endHeight
	}
}

func moveDown(bricks []oldbrick) ([]oldbrick, int) {
	moved := true
	movedCounter := 0
	for moved {
		moved = false
		var newBricks []oldbrick
		for _, this := range bricks {
			lowerZ := this.start[2]
			if this.end[2] < lowerZ {
				lowerZ = this.end[2]
			}
			afterFallHeight := lowerZ
			//for _,  := range bricks {

			//xStep := b.end[0] - b.start[0]
			//if utils.IntAbs(xStep) > 0 {
			//	xStep =
			//}
			//yStep := b.end[1] - b.start[1]
			//if utils.IntAbs(yStep) > 0 {
			//	yStep =
			//}
			if this.start[2] == 1 || this.end[2] == 1 {
				newBricks = append(newBricks, this)
				continue
			}
			if (this.end[0]-this.start[0]) == 0 && (this.end[1]-this.start[1]) == 0 {
				afterFallHeight = getFallDepth(this.start[0], this.end[1], lowerZ, bricks, this)
			} else {
				var xStep int
				var yStep int
				if this.end[0]-this.start[0] == 0 {
					xStep = 1
				} else {
					xStep = (this.end[0] - this.start[0]) / utils.IntAbs(this.end[0]-this.start[0])
				}
				if this.end[1]-this.start[1] == 0 {
					yStep = 1
				} else {
					yStep = (this.end[1] - this.start[1]) / utils.IntAbs(this.end[1]-this.start[1])
				}
				for x := this.start[0]; x <= this.end[0]; x += xStep {
					for y := this.start[1]; y <= this.end[1]; y += yStep {
						//for z := b.start[2]; z < b.end[2]; z += (b.end[2] - b.start[2]) / utils.IntAbs(b.end[2]-b.start[2]) {
						if aFH := getFallDepth(x, y, this.start[2], bricks, this); aFH < afterFallHeight {
							afterFallHeight = aFH
						}
						//}
					}
				}
			}
			//}
			newBrick := this
			if afterFallHeight != 0 {
				newBrick.start[2] -= afterFallHeight
				newBrick.end[2] -= afterFallHeight
				moved = true
				movedCounter++
			}
			newBricks = append(newBricks, newBrick)

		}
		bricks = newBricks
	}
	return bricks, movedCounter
}

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
outer:
	for i, _ := range bricks {
		brickCopy := slices.Clone(bricks)
		oneBrickRemoved := append(brickCopy[:i], brickCopy[(i+1):]...)
		for _, b := range oneBrickRemoved {
			if _, droped := dropBrick(b, oneBrickRemoved); droped {
				continue outer
			}
		}
		removeCounter++
	}
	fmt.Println(removeCounter)

}
