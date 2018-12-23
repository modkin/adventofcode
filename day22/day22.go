package day22

import "fmt"

func typeMap(depth int, x int, y int) [][]int {
	typeMap := make([][]int, x+1)
	for i := range typeMap {
		typeMap[i] = make([]int, y+1)
	}
	var geoIndex int
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			if i == 0 && j == 0 {
				geoIndex = 0
			} else if i == x && j == y {
				geoIndex = 0
			} else if j == 0 {
				geoIndex = i * 16807
			} else if i == 0 {
				geoIndex = j * 48271
			} else {
				geoIndex = typeMap[i-1][j] * typeMap[i][j-1]
			}
			erosionLevel := (geoIndex + depth) % 20183
			typeMap[i][j] = erosionLevel
		}
	}
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			typeMap[i][j] = typeMap[i][j] % 3
		}
	}
	return typeMap
}

func printTypeMap(typmap [][]int) {
	for y := 0; y < len(typmap[0]); y++ {
		for x := 0; x < len(typmap); x++ {
			switch typmap[x][y] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("=")
			case 2:
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func sumTypeMap(typemap [][]int) (ret int) {
	for y := 0; y < len(typemap[0]); y++ {
		for x := 0; x < len(typemap); x++ {
			ret += typemap[x][y]
		}
	}
	return
}

func Task1() {
	typemap := typeMap(4080, 14, 785)
	//printTypeMap(typemap)
	fmt.Println(sumTypeMap(typemap))
}
