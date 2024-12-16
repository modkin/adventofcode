package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type robot struct {
	pos [2]int
	dir [2]int
}

func add2Int(one, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

func moveRobot(robot robot, xMax, yMax int) [2]int {
	newPos := add2Int(robot.pos, robot.dir)
	if newPos[0] < 0 {
		newPos[0] = xMax - utils.IntAbs(newPos[0])
	}
	if newPos[0] >= xMax {
		newPos[0] = newPos[0] % xMax
	}
	if newPos[1] < 0 {
		newPos[1] = yMax - utils.IntAbs(newPos[1])
	}
	if newPos[1] >= yMax {
		newPos[1] = newPos[1] % yMax
	}
	return newPos
}

func paint(robots []robot, num, xMax, yMax int) {
	//fmt.Println("----------------", num, "---------------")
	allPos := make(map[[2]int]int)
	for _, r := range robots {
		allPos[r.pos]++
	}
	//utils.Print2DIntGrid(allPos)
	//fmt.Println("----------------", num, "---------------")
	file, err := os.Create("2024/day14/pics/" + strconv.Itoa(num) + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for y := 0; y <= yMax; y++ {
		line := ""
		for x := 0; x <= xMax; x++ {
			if _, ok := allPos[[2]int{x, y}]; ok {
				line += "."
			} else {
				line += " "
			}
		}
		file.WriteString(line + "\n")
	}
	file.Sync()
}

func check(robots []robot, xMax, yMax int) [4]int {
	allPos := make(map[[2]int]int)

	for _, r := range robots {

		allPos[r.pos]++
	}

	quad := [4]int{0, 0, 0, 0}

	for p, _ := range allPos {
		//p := r.pos
		if p[0] < xMax/2 && p[1] < yMax/2 {
			quad[0]++
		}
		if p[0] > xMax/2 && p[1] < yMax/2 {
			quad[1]++
		}
		if p[0] < xMax/2 && p[1] > yMax/2 {
			quad[2]++
		}
		if p[0] > xMax/2 && p[1] > yMax/2 {
			quad[3]++
		}
	}

	return quad
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day14/input")

	reg := regexp.MustCompile(`p=(-?\d*),(-?\d*) v=(-?\d*),(-?\d*)`)
	robots := []robot{}
	for _, line := range lines {

		tmp := reg.FindStringSubmatch(line)
		newRobot := robot{pos: [2]int{utils.ToInt(tmp[1]), utils.ToInt(tmp[2])}, dir: [2]int{utils.ToInt(tmp[3]), utils.ToInt(tmp[4])}}
		robots = append(robots, newRobot)
	}

	xMax := 101
	yMax := 103

	//for step := 0; step < 12; step++ {
	//	for i, _ := range robots {
	//		robots[i].pos = moveRobot(robots[i], xMax, yMax)
	//	}
	//}
	//paint(robots, 0)
	//for outer := 0; outer < 100; outer++ {
	//	for step := 0; step < 101; step++ {
	//		for i, _ := range robots {
	//			robots[i].pos = moveRobot(robots[i], xMax, yMax)
	//
	//		}
	//	}
	//	paint(robots, outer)
	//}

	counter := 0
	for {
		for i, _ := range robots {
			robots[i].pos = moveRobot(robots[i], xMax, yMax)
		}
		//quads := check(robots, xMax, yMax)
		//if quads[0] == quads[1] && quads[2] == quads[3] {
		//	fmt.Println(quads)
		//	paint(robots, 0)
		//
		//}
		counter++
		if (counter-12)%101 == 0 {
			paint(robots, counter, xMax, yMax)
		}
		if counter == 10000 {
			break
		}
	}

	prod := 1

	fmt.Println(robots)

	quad := [4]int{0, 0, 0, 0}

	for _, r := range robots {
		p := r.pos
		if p[0] < xMax/2 && p[1] < yMax/2 {
			quad[0]++
		}
		if p[0] > xMax/2 && p[1] < yMax/2 {
			quad[1]++
		}
		if p[0] < xMax/2 && p[1] > yMax/2 {
			quad[2]++
		}
		if p[0] > xMax/2 && p[1] > yMax/2 {
			quad[3]++
		}
	}

	fmt.Println(quad)
	fmt.Println(quad[0] * quad[1] * quad[2] * quad[3])

	fmt.Println(prod)
}
