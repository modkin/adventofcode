package day11

import "fmt"

func calcPower(grid *[300][300]int, sn int) {
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			rackID := x + 1 + 10
			powerlvl := rackID * (1 + y)
			powerlvl = powerlvl + sn
			powerlvl = powerlvl * rackID
			powerlvl = powerlvl % 1000
			powerlvl = int(powerlvl / 100)
			powerlvl -= 5
			grid[x][y] = powerlvl
		}
	}
}

func sumPower(grid *[300][300]int, xSize int, ySize int) (outgrid [300][300]int) {
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			for i := 0; i < xSize; i++ {
				for j := 0; j < ySize; j++ {
					if x+i < 300 && y+j < 300 {
						outgrid[x][y] += grid[x+i][y+j]
					}
				}
			}
		}
	}
	return
}

func findMax(grid *[300][300]int) (max, xout, yout int) {
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			if grid[x][y] > max {
				max = grid[x][y]
				xout = x + 1
				yout = y + 1
			}
		}
	}
	return
}

func findMaxAnySize(grid *[300][300]int) (xout, yout, size int) {
	max := 0
	for i := 0; i < 75; i++ {
		sumGrid := sumPower(grid, i, i)
		currmax, x, y := findMax(&sumGrid)
		if currmax > max {
			max = currmax
			xout = x
			yout = y
			size = i
		}
	}
	return
}

func doTests() {
	var grid [300][300]int
	calcPower(&grid, 8)
	fmt.Println(grid[2][4])
	calcPower(&grid, 57)
	fmt.Println(grid[121][78])
	calcPower(&grid, 39)
	fmt.Println(grid[216][195])
	calcPower(&grid, 71)
	fmt.Println(grid[100][152])
	calcPower(&grid, 42)
	sumGrid := sumPower(&grid, 3, 3)
	_, x, y := findMax(&sumGrid)
	fmt.Println("expected 21,61: ", x, y)
	fmt.Println("expected 30 ", sumGrid[x-1][y-1])
	calcPower(&grid, 18)
	sumGrid = sumPower(&grid, 3, 3)
	_, x, y = findMax(&sumGrid)
	fmt.Println("expected 33,45: ", x, y)
	fmt.Println("expected 29 ", sumGrid[x-1][y-1])

	calcPower(&grid, 18)
	x, y, size := findMaxAnySize(&grid)
	fmt.Println("expected 90,269,16: ", x, y, size)
}

func Task1() {
	//doTests()
	var grid [300][300]int
	calcPower(&grid, 2187)
	sumGrid := sumPower(&grid, 3, 3)
	fmt.Println(findMax(&sumGrid))
	fmt.Println(findMaxAnySize(&grid))
}
