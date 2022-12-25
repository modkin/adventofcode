package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

func turnLeft(vec [2]int) (ret [2]int) {
	ret[0] = vec[1]
	ret[1] = -1 * vec[0]
	return
}

func turnRight(vec [2]int) (ret [2]int) {
	ret[0] = -1 * vec[1]
	ret[1] = vec[0]
	return
}

func colMinMax(grid map[[2]int]string, x int) (min, max int) {
	outside := true
	for y := 0; y < math.MaxInt; y++ {
		if outside {
			if _, ok := grid[[2]int{x, y}]; ok {
				min = y
				outside = false
			}
		} else {
			if _, ok := grid[[2]int{x, y}]; !ok {
				max = y - 1
				return
			}
		}
	}
	fmt.Println("ERROR")
	return -math.MaxInt, math.MaxInt
}

func rowMinMax(grid map[[2]int]string, y int) (min, max int) {
	outside := true
	for x := 0; x < math.MaxInt; x++ {
		if outside {
			if _, ok := grid[[2]int{x, y}]; ok {
				min = x
				outside = false
			}
		} else {
			if _, ok := grid[[2]int{x, y}]; !ok {
				max = x - 1
				return
			}
		}
	}
	fmt.Println("ERROR")
	return -math.MaxInt, math.MaxInt
}

func add(first [2]int, second [2]int) [2]int {
	return [2]int{first[0] + second[0], first[1] + second[1]}
}

func mul(first [2]int, prod int) [2]int {
	return [2]int{first[0] * prod, first[1] * prod}
}

func sub(first [2]int, second [2]int) [2]int {
	return [2]int{first[0] - second[0], first[1] - second[1]}
}

func dist(first [2]int, second [2]int) int {
	ret := 0
	ret += utils.IntAbs(second[0] - first[0])
	ret += utils.IntAbs(second[1] - first[1])
	return ret
}

func norm(in [2]int) [2]int {
	if in[0] == 0 {
		return [2]int{0, in[1] / utils.IntAbs(in[1])}
	} else if in[1] == 0 {
		return [2]int{in[0] / utils.IntAbs(in[0]), 0}
	} else {
		fmt.Println("ERROR")
	}
	return [2]int{0, 0}
}

func move(grid map[[2]int]string, start [2]int, dir [2]int, steps int) (pos [2]int) {
	pos = start
	for i := 0; i < steps; i++ {
		if tile, ok := grid[add(pos, dir)]; ok {
			if tile == "#" {
				return
			} else {
				pos = add(pos, dir)
			}
		} else {
			if dir[0] != 0 {
				min, max := rowMinMax(grid, pos[1])
				if dir[0] == 1 {
					if tar := [2]int{min, pos[1]}; grid[tar] == "#" {
						return
					} else {
						pos = tar
					}
				} else if dir[0] == -1 {
					if tar := [2]int{max, pos[1]}; grid[tar] == "#" {
						return
					} else {
						pos = tar
					}
				}
			} else {
				min, max := colMinMax(grid, pos[0])
				if dir[1] == 1 {
					if tar := [2]int{pos[0], min}; grid[tar] == "#" {
						return
					} else {
						pos = tar
					}
				} else if dir[1] == -1 {
					if tar := [2]int{pos[0], max}; grid[tar] == "#" {
						return
					} else {
						pos = tar
					}
				}
			}
		}
	}
	return pos

}

func inBetween(pos, start, end [2]int) bool {
	if pos[0] == start[0] && pos[0] == end[0] {
		tmp := []int{start[1], end[1]}
		sort.Ints(tmp)
		if pos[1] >= tmp[0] && pos[1] <= tmp[1] {
			return true
		}
	}
	if pos[1] == start[1] && pos[1] == end[1] {
		tmp := []int{start[0], end[0]}
		sort.Ints(tmp)
		if pos[0] >= tmp[0] && pos[0] <= tmp[1] {
			return true
		}
	}
	return false
}

func getInwardDir(edge [2][2]int, sideSize int) [2]int {
	if edge[0][0] == edge[1][0] {
		if edge[0][0]%sideSize == 0 {
			return [2]int{1, 0}
		} else if edge[0][0]%sideSize == sideSize-1 {
			return [2]int{-1, 0}
		} else {
			fmt.Println("ERROR 181")
		}
	} else if edge[0][1] == edge[1][1] {
		if edge[1][1]%sideSize == 0 {
			return [2]int{0, 1}
		} else if edge[1][1]%sideSize == sideSize-1 {
			return [2]int{0, -1}
		} else {
			fmt.Println("ERROR 189")
		}
	}
	fmt.Println("ERROR 192")
	return [2]int{math.MaxInt, math.MaxInt}
}

func move2(grid map[[2]int]string, start [2]int, dir [2]int, steps int, edgeMap map[[2]int][][2][2]int, sideSize int) ([2]int, [2]int) {

	pos := start
	for i := 0; i < steps; i++ {
		if tile, ok := grid[add(pos, dir)]; ok {
			if tile == "#" {
				return pos, dir
			} else {
				pos = add(pos, dir)
			}
		} else {
		findEdge:
			for _, edges := range edgeMap {
				for idx, edge := range edges {
					if inBetween(pos, edge[0], edge[1]) {
						if edge[0][0] == edge[1][0] {
							if dir[1] != 0 {
								continue
							}
						} else if edge[0][1] == edge[1][1] {
							if dir[0] != 0 {
								continue
							}
						}
						//tmp := sub(pos, edge[0])
						stepsFromCorner := dist(edge[0], pos)
						edgeVect := sub(edges[1-idx][1], edges[1-idx][0])
						tar := add(edges[1-idx][0], mul(norm(edgeVect), stepsFromCorner))
						if grid[tar] == "#" {
							return pos, dir
						} else {
							pos = tar
							dir = getInwardDir(edges[1-idx], sideSize)
						}
						break findEdge
					}
				}
			}
		}
	}
	return pos, dir

}

func addSide(edgeMapping map[[2]int][][2][2]int, start [2]int, dir [2]int, corners [4]int, sideSize int) {
	for i := 0; i <= 3; i++ {
		stop := add(start, mul(dir, sideSize-1))
		edge := [2]int{corners[i], corners[(i+1)%4]}
		edgeTurned := [2]int{corners[(i+1)%4], corners[i]}
		if val, ok := edgeMapping[edge]; ok {
			edgeMapping[edge] = append(val, [2][2]int{start, stop})
		} else if val, ok = edgeMapping[edgeTurned]; ok {
			edgeMapping[edgeTurned] = append(val, [2][2]int{stop, start})
		} else {
			edgeMapping[edge] = append(edgeMapping[edge], [2][2]int{start, stop})
		}
		start = stop
		dir = turnRight(dir)
	}
}

func main() {
	const filename = "2022/day22/input"
	var sideSize int
	if strings.Contains(filename, "test") {
		sideSize = 4
	} else {
		sideSize = 50
	}
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	grid := make(map[[2]int]string)
	path := make([]string, 0)
	//sideMappings := make(map[[sideSize][2]int][sideSize][2]int)
	edgeMapping := make(map[[2]int][][2][2]int)

	yPos := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			scanner.Scan()
			line := scanner.Text()
			numreg := regexp.MustCompile(`\d*`)
			dirreg := regexp.MustCompile(`[RL]`)
			numbers := numreg.FindAllString(line, -1)
			dirs := dirreg.FindAllString(line, -1)
			for i, dir := range dirs {
				path = append(path, numbers[i])
				path = append(path, dir)
			}
			path = append(path, numbers[len(numbers)-1])
		} else {

			line := strings.Split(scanner.Text(), "")
			for xPos, s := range line {
				if s != " " {
					grid[[2]int{xPos, yPos}] = s
				}
			}
			yPos++
		}
	}

	dir := [2]int{1, 0}
	if sideSize == 4 {
		addSide(edgeMapping, [2]int{8, 0}, dir, [4]int{0, 1, 2, 3}, sideSize)
		addSide(edgeMapping, [2]int{8, 4}, dir, [4]int{3, 2, 6, 7}, sideSize)
		addSide(edgeMapping, [2]int{4, 4}, dir, [4]int{0, 3, 7, 4}, sideSize)
		addSide(edgeMapping, [2]int{0, 4}, dir, [4]int{1, 0, 4, 5}, sideSize)
		addSide(edgeMapping, [2]int{8, 8}, dir, [4]int{7, 6, 5, 4}, sideSize)
		addSide(edgeMapping, [2]int{12, 8}, dir, [4]int{6, 2, 1, 5}, sideSize)
	} else {
		addSide(edgeMapping, [2]int{50, 0}, dir, [4]int{0, 1, 2, 3}, sideSize)
		addSide(edgeMapping, [2]int{100, 0}, dir, [4]int{1, 5, 6, 2}, sideSize)
		addSide(edgeMapping, [2]int{50, 50}, dir, [4]int{3, 2, 6, 7}, sideSize)
		addSide(edgeMapping, [2]int{50, 100}, dir, [4]int{7, 6, 5, 4}, sideSize)
		addSide(edgeMapping, [2]int{0, 100}, dir, [4]int{3, 7, 4, 0}, sideSize)
		addSide(edgeMapping, [2]int{0, 150}, dir, [4]int{0, 4, 5, 1}, sideSize)
	}

	fmt.Println("laenge:", len(edgeMapping))
	for ints, i := range edgeMapping {
		fmt.Println(ints)
		fmt.Println(i)
		fmt.Println("NEXT")
	}

	//fmt.Println(path)
	//utils.Print2DStringsGrid(grid)
	//fmt.Println(move2(grid, [2]int{6, 4}, [2]int{0, -1}, 1, edgeMapping, sideSize))

	startx, _ := rowMinMax(grid, 0)
	pos := [2]int{startx, 0}
	dir = [2]int{1, 0}
	for i, s := range path {
		if i%2 == 0 {
			steps := utils.ToInt(s)
			pos, dir = move2(grid, pos, dir, steps, edgeMapping, sideSize)
		} else {
			if s == "L" {
				dir = turnLeft(dir)
			} else {
				dir = turnRight(dir)
			}
			grid[pos] = "*"
			//utils.Print2DStringsGrid(grid)
		}
	}
	//utils.Print2DStringsGrid(grid)
	row := pos[1] + 1
	col := pos[0] + 1
	password := row*1000 + col*4
	if dir == [2]int{1, 0} {
		password += 0
	} else if dir == [2]int{0, 1} {
		password += 1
	} else if dir == [2]int{-1, 0} {
		password += 2
	} else if dir == [2]int{0, -1} {
		password += 3
	}
	fmt.Println("Day 22:", password)
}
