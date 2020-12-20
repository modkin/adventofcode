package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

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

func findFourNeighbors(tiles [][10][10]string, tile *[10][10]string) bool {
	top, bot, left, right := false, false, false, false
	for _, elem := range tiles {
		if elem == *tile {
			continue
		}
		//top := tile[:][0]
		//bot := tile[:][9]
		//left := tile[0][:]
		//rigt := tile[9][:]
		for _, otherTile := range getVariants(&elem) {
			top = (tile[:][0] == otherTile[:][0]) || top
			bot = (tile[:][9] == otherTile[:][9]) || bot
			left = sameStringSlice(tile[0][:], otherTile[0][:]) || left
			right = sameStringSlice(tile[9][:], otherTile[9][:]) || right
		}
	}
	if top && bot && left && right {
		return true
	} else {
		return false
	}
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day20/testinput"))
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
		if !findFourNeighbors(tiles, &tile) {
			mult *= ids[i]
			fmt.Println(i)
		}
	}

	//for _, tile := range getVariants(&tiles[0]) {
	//	utils.Print1010Grid(tile)
	//}
}
