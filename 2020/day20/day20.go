package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func rotateTile90(tile *[10][10]string) [10][10]string {
	var rotate90 [10][10]string
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			rotate90[x][y] = tile[y][9-x]
		}
	}
	return rotate90
}

func getAllVariants(tile *[10][10]string) [][10][10]string {
	var rotate0 [10][10]string
	var flipH [10][10]string
	var flipV [10][10]string
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			rotate0[x][y] = tile[x][y]
			flipH[x][y] = tile[x][9-y]
			flipV[x][y] = tile[9-x][y]
		}
	}
	ret := [][10][10]string{rotate0, flipH, flipV}
	for i := 0; i <= 2; i++ {
		ret = append(ret, rotateTile90(&ret[0+i*3]))
		ret = append(ret, rotateTile90(&ret[1+i*3]))
		ret = append(ret, rotateTile90(&ret[2+i*3]))
	}
	return ret
}

func getVariants(tile *[10][10]string) [][10][10]string {
	//ret := make([][10][10]string,0)
	var rotate0 [10][10]string
	var rotate90 [10][10]string
	var rotate180 [10][10]string
	var rotate270 [10][10]string
	var flipH [10][10]string
	var flipV [10][10]string
	//rotate
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			rotate0[x][y] = tile[x][y]
			rotate90[x][y] = tile[y][9-x]
			rotate180[x][y] = tile[9-x][9-y]
			rotate270[x][y] = tile[9-y][x]
			flipH[x][y] = tile[x][9-y]
			flipV[x][y] = tile[9-x][y]
		}
	}
	return [][10][10]string{rotate0, rotate90, rotate180, rotate270, flipH, flipV}
}

func sameStringSlice(one []string, two []string) bool {
	same := true
	for i, elem := range one {
		if elem != two[i] {
			same = false
		}
	}
	return same
}

func possibleNeighbors(tiles [][10][10]string, inputTile *[10][10]string, ids []int) [][]int {
	neighborsAtOrientation := make([][]int, 0)
	neighborsFound := 0
	for _, tile := range getAllVariants(inputTile) {
		neighborsSlice := make([]int, 0)
		for otherId, elem := range tiles {
			//fmt.Println("Other Tile:", ids[i])
			if elem == tile {
				continue
			}
			for _, otherTile := range getAllVariants(&elem) {
				//top
				if tile[:][0] == otherTile[:][0] {
					neighborsSlice = append(neighborsSlice, ids[otherId])
					neighborsFound++
				}
				//bot
				if tile[:][9] == otherTile[:][9] {
					neighborsSlice = append(neighborsSlice, ids[otherId])
					neighborsFound++
				}
				//left
				if sameStringSlice(tile[0][:], otherTile[0][:]) {
					neighborsSlice = append(neighborsSlice, ids[otherId])
					neighborsFound++
				}
				//right
				if sameStringSlice(tile[9][:], otherTile[9][:]) {
					neighborsSlice = append(neighborsSlice, ids[otherId])
					neighborsFound++
				}
			}
		}
		neighborsAtOrientation = append(neighborsAtOrientation, neighborsSlice)
	}
	return neighborsAtOrientation
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day20/input"))
	tiles := make([][10][10]string, 0)
	ids := make([]int, 0)
	var newTile [10][10]string
	y := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			tiles = append(tiles, newTile)
			newTile = [10][10]string{}
			y = 0
		} else if strings.Contains(scanner.Text(), "Tile") {
			tmp := strings.Trim(strings.Split(scanner.Text(), " ")[1], ":")
			ids = append(ids, utils.ToInt(tmp))
		} else {
			for x, elem := range strings.Split(scanner.Text(), "") {
				newTile[x][y] = elem
			}
			y++
		}
	}
	//add last tile
	tiles = append(tiles, newTile)

	mult := 1
	for i, tile := range tiles {
		//fmt.Println("Tile:,",ids[i])
		tmp := possibleNeighbors(tiles, &tile, ids)
		max := 0
		for _, nbrs := range tmp {
			tmpMap := make(map[int]bool)
			for _, elem := range nbrs {
				tmpMap[elem] = true
			}
			if len(tmpMap) > max {
				max = len(tmpMap)
			}
		}
		if max == 2 {
			mult *= ids[i]
		}
	}
	fmt.Println(mult)

	//mult := 1
	//for i, tile := range tiles {
	//	if possibleNeighbors(tiles, &tile, ids) {
	//		mult *= ids[i]
	//		fmt.Println("True:",ids[i])
	//	} else {
	//		//fmt.Println("False:", ids[i])
	//	}
	//}

	//for _, tile := range getVariants(&tiles[0]) {
	//	utils.Print1010Grid(tile)
	//}
}
