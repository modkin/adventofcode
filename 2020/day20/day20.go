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

func checkTopBot(upper [10][10]string, lower [10][10]string) bool {
	for i := 0; i < 10; i++ {
		if upper[i][9] != lower[i][0] {
			return false
		}
	}
	return true
}

func checkLeftRight(left [10][10]string, right [10][10]string) bool {
	for i := 0; i < 10; i++ {
		if left[9][i] != right[0][i] {
			return false
		}
	}
	return true
}

func findSeaMonster(image [96][96]string) int {
	MonstersFound := 0
	for y := 0; y < 96-2; y++ {
		for x := 0; x < 96-19; x++ {
			found := make([]string, 0)
			found = append(found, image[x+18][y])

			found = append(found, image[x][y+1])
			found = append(found, image[x+5][y+1])
			found = append(found, image[x+6][y+1])
			found = append(found, image[x+11][y+1])
			found = append(found, image[x+12][y+1])
			found = append(found, image[x+17][y+1])
			found = append(found, image[x+18][y+1])
			found = append(found, image[x+19][y+1])

			found = append(found, image[x+1][y+2])
			found = append(found, image[x+4][y+2])
			found = append(found, image[x+7][y+2])
			found = append(found, image[x+10][y+2])
			found = append(found, image[x+13][y+2])
			found = append(found, image[x+16][y+2])
			if strings.Count(strings.Join(found, ""), "#") == 15 {
				MonstersFound++
			}
		}
	}
	return MonstersFound
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day20/input"))
	allTiles := make(map[int][10][10]string)
	var newTile [10][10]string
	y := 0
	newTileId := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			allTiles[newTileId] = newTile
			newTile = [10][10]string{}
			y = 0
		} else if strings.Contains(scanner.Text(), "Tile") {
			tmp := strings.Trim(strings.Split(scanner.Text(), " ")[1], ":")
			newTileId = utils.ToInt(tmp)
		} else {
			for x, elem := range strings.Split(scanner.Text(), "") {
				newTile[x][y] = elem
			}
			y++
		}
	}
	neighbors := make(map[int][]int)
	for tileId, tile := range allTiles {
		for otherTileId, otherTile := range allTiles {
			if otherTileId == tileId {
				continue
			}

			for _, rotation := range getAllVariants(&otherTile) {
				if checkTopBot(tile, rotation) {
					if !utils.IntSliceContains(neighbors[tileId], otherTileId) {
						neighbors[tileId] = append(neighbors[tileId], otherTileId)
					}
				}
				if checkTopBot(rotation, tile) {
					if !utils.IntSliceContains(neighbors[tileId], otherTileId) {
						neighbors[tileId] = append(neighbors[tileId], otherTileId)
					}
				}
				if checkLeftRight(rotation, tile) {
					if !utils.IntSliceContains(neighbors[tileId], otherTileId) {
						neighbors[tileId] = append(neighbors[tileId], otherTileId)
					}
				}
				if checkLeftRight(tile, rotation) {
					if !utils.IntSliceContains(neighbors[tileId], otherTileId) {
						neighbors[tileId] = append(neighbors[tileId], otherTileId)
					}
				}
			}
		}
	}
	cornerIds, edgeIds := make([]int, 0), make([]int, 0)
	for tileId, neighborIds := range neighbors {
		if len(neighborIds) == 2 {
			cornerIds = append(cornerIds, tileId)
		} else if len(neighborIds) == 3 {
			edgeIds = append(edgeIds, tileId)
		}
	}
	solutionOne := cornerIds[0] * cornerIds[1] * cornerIds[2] * cornerIds[3]
	fmt.Println("Task 20.1:", solutionOne)
	//fmt.Println("Corners: ", cornerIds)

	var fullPicture [12][12]int

	notInPicture := func(id int) bool {
		for _, line := range fullPicture {
			for _, elem := range line {
				if elem == id {
					return false
				}
			}
		}
		return true
	}
	startId := 1439
	fullPicture[0][0] = startId
	fullPicture[1][0] = 2417
	fullPicture[0][1] = 3461
	for i := 1; i < 12; i++ {
		for _, nbr := range neighbors[fullPicture[0][i]] {
			if utils.IntSliceContains(edgeIds, nbr) && notInPicture(nbr) {
				fullPicture[0][i+1] = nbr
			}
		}
		for _, nbr := range neighbors[fullPicture[i][0]] {
			if utils.IntSliceContains(edgeIds, nbr) && notInPicture(nbr) {
				fullPicture[i+1][0] = nbr
			}
		}
	}
	for _, nbr := range neighbors[fullPicture[0][10]] {
		if utils.IntSliceContains(cornerIds, nbr) {
			fullPicture[0][11] = nbr
		}
	}
	for _, nbr := range neighbors[fullPicture[10][0]] {
		if utils.IntSliceContains(cornerIds, nbr) {
			fullPicture[11][0] = nbr
		}
	}
	for x := 1; x < 12; x++ {
		for y = 1; y < 12; y++ {
			for _, nbr := range neighbors[fullPicture[x][y-1]] {
				if utils.IntSliceContains(neighbors[fullPicture[x-1][y]], nbr) && notInPicture(nbr) {
					fullPicture[x][y] = nbr
				}

			}
		}
	}
	//fmt.Println(len(edgeIds))
	//fmt.Println(len(neighbors))

	flatFullPicture := make([][][10][10]string, 12)
	for i := 0; i < 12; i++ {
		flatFullPicture[i] = make([][10][10]string, 12)
		for foo := 0; foo < 12; foo++ {
			var tmp [10][10]string
			for a := 0; a < 10; a++ {
				for b := 0; b < 10; b++ {
					tmp[a][b] = "-"
				}
			}
			flatFullPicture[i][foo] = tmp
		}
	}
	tmp := allTiles[1439]
	flatFullPicture[0][0] = getAllVariants(&tmp)[7]
	//outer:
	for x := 0; x < 12; x++ {
		if x != 11 {
			currentTile := flatFullPicture[x][0]
			nbrTmp := allTiles[fullPicture[x+1][0]]
			for _, nbrTile := range getAllVariants(&nbrTmp) {
				if checkLeftRight(currentTile, nbrTile) {
					flatFullPicture[x+1][0] = nbrTile
				}
			}
		}
		for y = 0; y < 11; y++ {
			currentTile := flatFullPicture[x][y]
			nbrTmp := allTiles[fullPicture[x][y+1]]
			for _, nbrTile := range getAllVariants(&nbrTmp) {
				if checkTopBot(currentTile, nbrTile) {
					flatFullPicture[x][y+1] = nbrTile
				}
			}
		}
	}
	//for yTile := 0; yTile < 12; yTile++ {
	//	for y = 0; y < 10; y++ {
	//		for tileNr := 0; tileNr < 12; tileNr++ {
	//			for x := 0; x < 10; x++ {
	//				fmt.Print(flatFullPicture[tileNr][yTile][x][y])
	//			}
	//			fmt.Print(" ")
	//		}
	//		fmt.Println()
	//	}
	//	fmt.Println()
	//}
	var finnalImage [96][96]string
	for yTile := 0; yTile < 12; yTile++ {
		for y = 1; y < 9; y++ {
			for tileNr := 0; tileNr < 12; tileNr++ {
				for x := 1; x < 9; x++ {
					finnalImage[(x-1)+tileNr*8][(y-1)+yTile*8] = flatFullPicture[tileNr][yTile][x][y]
				}
			}
		}
	}
	//for y := 0; y < 96; y++ {
	//	for x := 0; x < 96; x++ {
	//		fmt.Print(finnalImage[x][y])
	//	}
	//	fmt.Println()
	//}
	numberOfMonsters := findSeaMonster(finnalImage)
	countHash := 0
	for y := 0; y < 96; y++ {
		for x := 0; x < 96; x++ {
			if finnalImage[x][y] == "#" {
				countHash++
			}
		}
	}
	fmt.Println("Task 20.2:", countHash-15*numberOfMonsters)
}
