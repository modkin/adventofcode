package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func CopyIntSlice(input []int) []int {
	sliceCopy := make([]int, len(input))
	copy(sliceCopy, input)
	return sliceCopy
}

func CopyStringSlice(input []string) []string {
	sliceCopy := make([]string, len(input))
	copy(sliceCopy, input)
	return sliceCopy
}

func ToInt64(str string) int64 {
	ret, err := strconv.ParseInt(strings.TrimSpace(str), 10, 64)
	if err != nil {
		panic(err)
	}
	return int64(ret)
}

func ToInt(str string) int {
	ret, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		panic(err)
	}
	return ret
}

func IntAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func Int64Abs(x int64) int64 {
	return int64(math.Abs(float64(x)))
}

func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func SumSlice(slice []int) int {
	sum := 0
	for _, element := range slice {
		sum += element
	}
	return sum
}

func ReverseSlice(a []string) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func ReverseIntSlice(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

// /split integer into integer slice of the digits
func SplitInt(input int) []int {
	count := len(fmt.Sprint(input))
	output := make([]int, count)
	for i := 0; i < count; i++ {
		output[i] = input / int(math.Pow(10, float64(count-1-i))) % 10
	}
	return output
}

func MaxIdx(toCheck [2]int) int {
	if IntAbs(toCheck[0]) > IntAbs(toCheck[1]) {
		return 0
	} else {
		return 1
	}
}

func Sum(one [2]int, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

func SliceContains(list []string, s string) bool {
	for _, elem := range list {
		if s == elem {
			return true
		}
	}
	return false
}

func RuneSliceContains(list []rune, s rune) bool {
	for _, elem := range list {
		if s == elem {
			return true
		}
	}
	return false
}

func CountStringinStringSlice(list []string, s string) int {
	count := 0
	for _, elem := range list {
		if s == elem {
			count++
		}
	}
	return count
}

func IntSliceContains(list []int, s int) bool {
	for _, elem := range list {
		if s == elem {
			return true
		}
	}
	return false
}

func Int2SliceContains(list [][2]int, s [2]int) bool {
	for _, elem := range list {
		if s == elem {
			return true
		}
	}
	return false
}

func Print1010Grid(mapToPrint [10][10]string) {
	fmt.Println("=============")
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Print(mapToPrint[x][y])
		}
		fmt.Println()
	}
	fmt.Println("=============")
}

func Print3DGrid(grid map[[3]int]bool, max int) {
	for z := -max; z < max; z++ {
		fmt.Println("z=", z)
		for y := -max; y < max+2; y++ {
			for x := -max; x < max+2; x++ {
				if grid[[3]int{x, y, z}] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}

	}
}

func Print2DIntGrid(grid map[[2]int]int) {
	xMax, yMax := 0, 0
	for i := range grid {
		if i[0] > xMax {
			xMax = i[0]
		}
		if i[1] > yMax {
			yMax = i[1]
		}
	}
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Print2DStringsGrid(grid map[[2]int]string) {
	fmt.Println("----------------------------------------------")
	xMin, yMin, xMax, yMax := math.MaxInt, math.MaxInt, 0, 0
	for i := range grid {
		if i[0] < xMin {
			xMin = i[0]
		}
		if i[0] > xMax {
			xMax = i[0]
		}
		if i[1] < yMin {
			yMin = i[1]
		}
		if i[1] > yMax {
			yMax = i[1]
		}
	}
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Print2D2IntGrid(grid [][2]int) {
	fmt.Println("----------------------------------------------")
	xMin, yMin, xMax, yMax := math.MaxInt, math.MaxInt, 0, 0
	for _, i := range grid {
		if i[0] < xMin {
			xMin = i[0]
		}
		if i[0] > xMax {
			xMax = i[0]
		}
		if i[1] < yMin {
			yMin = i[1]
		}
		if i[1] > yMax {
			yMax = i[1]
		}
	}
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if Int2SliceContains(grid, [2]int{x, y}) {

				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Print2DStringGrid(grid map[[2]int]bool) {
	fmt.Println("----------------------------------------------")
	xMax, yMax := 0, 0
	for i := range grid {
		if i[0] > xMax {
			xMax = i[0]
		}
		if i[1] > yMax {
			yMax = i[1]
		}
	}
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if _, ok := grid[[2]int{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------------")
}

func PrintTetris(ingrid map[[2]int]string, shape [][2]int, xDiff, yDiff int) {
	grid := make(map[[2]int]string)
	for i, s := range ingrid {
		grid[i] = s
	}
	for _, i := range shape {
		grid[[2]int{i[0] + xDiff, i[1] + yDiff}] = "*"
	}
	fmt.Println("----------------------------------------------")
	xMax, yMax := 0, 0
	for i := range grid {
		if i[0] > xMax {
			xMax = i[0]
		}
		if i[1] > yMax {
			yMax = i[1]
		}
	}
	for y := yMax; y >= -1; y-- {
		for x := 0; x <= xMax; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------------")
}

func Factorial(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result = result * i
	}
	return result
}

func Sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
