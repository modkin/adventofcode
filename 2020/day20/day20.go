package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func printRow(row [][10][10]string) {
	for y := 0; y < 10; y++ {
		for tileNr := 0; tileNr < 12; tileNr++ {
			for x := 0; x < 10; x++ {
				fmt.Print(row[tileNr][x][y])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

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

func findPossibleNextLeft(current [10][10]string, tiles [][10][10]string, alreadyUsed []int) []int {
	ret := make([]int, 0)
	for i, nextTile := range tiles {
		if utils.IntSliceContains(alreadyUsed, i) {
			continue
		}
		for _, rotatedNexTile := range getAllVariants(&nextTile) {
			if sameStringSlice(current[9][:], rotatedNexTile[0][:]) {
				ret = append(ret, i)
			}
		}
	}
	return ret
}

func checkTopBot(upper [10][10]string, lower [10][10]string) bool {
	for i := 0; i < 10; i++ {
		if upper[i][9] != lower[i][0] {
			return false
		}
	}
	return true
}

func createRow(startId int, tiles [][10][10]string, alreadyUsed []int, currentRow []int, ret *[][]int, rowLength int, previousRow [][10][10]string) {
	currentRow = append(currentRow, startId)
	if len(currentRow) == rowLength {
		*ret = append(*ret, utils.CopyIntSlice(currentRow))
	} else {
		for _, rotation := range getAllVariants(&tiles[startId]) {
			if !checkTopBot(previousRow[len(currentRow)-1], rotation) {
				continue
			}
			possibleNextLeft := findPossibleNextLeft(rotation, tiles, alreadyUsed)
			if len(possibleNextLeft) == 0 {
				continue
			} else {
				possibleNextLeftMap := make(map[int]bool)
				for _, nextID := range possibleNextLeft {
					possibleNextLeftMap[nextID] = true
				}
				for key := range possibleNextLeftMap {
					alreadyUsed = append(alreadyUsed, key)
					createRow(key, tiles, utils.CopyIntSlice(alreadyUsed), utils.CopyIntSlice(currentRow), ret, rowLength, previousRow)
				}
			}
		}
	}
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
	//tiles = append(tiles, newTile)
	//fmt.Println(len(tiles))

	possibleStarts := make([][10][10]string, 0)
	cornerIdx := make([]int, 0)
	//mult := 1
	//for i, tile := range tiles {
	//	//fmt.Println("Tile:,",ids[i])
	//	tmp := possibleNeighbors(tiles, &tile, ids)
	//	max := 0
	//	for _, nbrs := range tmp {
	//		tmpMap := make(map[int]bool)
	//		for _, elem := range nbrs {
	//			tmpMap[elem] = true
	//		}
	//		if len(tmpMap) > max {
	//			max = len(tmpMap)
	//		}
	//	}
	//	if max == 2 {
	//		possibleStarts = append(possibleStarts, tile)
	//		cornerIdx = append(cornerIdx, i)
	//		mult *= ids[i]
	//	}
	//}
	cornerIdx = []int{10, 48, 64, 69}
	//fmt.Println("Task 1:",mult)
	if len(possibleStarts) != 4 || len(cornerIdx) != 4 {
		//fmt.Println("ERROR")
	}

	fmt.Println("Corners:", cornerIdx)

	//alreadyUsed := make([]int, 4)
	//copy(alreadyUsed, cornerIdx)

	firstRow := []int{10, 101, 29, 108, 97, 73, 2, 26, 96, 67, 80, 64}
	//possibleFirstRow := make([][]int, 0)

	firstRowTiles := make([][][10][10]string, 8)
	for firstRowCount := 0; firstRowCount < 8; firstRowCount++ {
		skip := 0
		for i, currentTileID := range firstRow {
		curRot:
			for _, current := range getAllVariants(&tiles[currentTileID]) {
				if i == 0 {
					firstRowTiles[firstRowCount] = [][10][10]string{current}
					if skip < firstRowCount {
						skip++
						continue curRot

					}

				}
				for _, rotation := range getAllVariants(&tiles[firstRow[i+1]]) {
					if sameStringSlice(current[9][:], rotation[0][:]) {
						firstRowTiles[firstRowCount] = append(firstRowTiles[firstRowCount], rotation)
						break curRot
					}
				}
			}
			if i == len(firstRow)-2 {
				break
			}
		}
	}
	fmt.Println(len(firstRowTiles))
	//for _, elem := range firstRowTiles {
	//	printRow(elem)
	//	fmt.Println("================================================================================")
	//}
	for _, toptile := range getAllVariants(&firstRowTiles[0][2]) {
		for j, bottile := range tiles {
			if checkTopBot(toptile, bottile) {
				fmt.Println("muh", j)
			}
		}
	}

}
