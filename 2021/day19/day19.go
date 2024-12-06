package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type conversion struct {
	offset [3]int
	rotMat [3][3]int
}

type beaconPair struct {
	beacon1  [3]int
	beacon2  [3]int
	distance float64
}

func getNormalizedDirection(start [3]int, end [3]int) [3]int {
	dir := [3]int{utils.IntAbs(end[0] - start[0]), utils.IntAbs(end[1] - start[1]), utils.IntAbs(end[2] - start[2])}
	//if dir[0] < 0 {
	//	dir[0] = -dir[0]
	//	dir[1] = -dir[1]
	//	dir[2] = -dir[2]
	//}
	return dir
}

func multVec(u, v [3]int) [3]int {
	return [3]int{u[0] * v[0], u[1] * v[1], u[2] * v[2]}
}

func addVec(u, v [3]int) [3]int {
	return [3]int{u[0] + v[0], u[1] + v[1], u[2] + v[2]}
}

func subVec(u, v [3]int) [3]int {
	return [3]int{u[0] - v[0], u[1] - v[1], u[2] - v[2]}
}

func divVec(u, v [3]int) [3]int {
	return [3]int{u[0] / v[0], u[1] / v[1], u[2] / v[2]}
}

func getDirection(start [3]int, end [3]int) [3]int {
	return [3]int{end[0] - start[0], end[1] - start[1], end[2] - start[2]}
}

func getAllDirs(scan [][3]int) [][3][3]int {
	var scanDirs [][3][3]int
	for i, beac := range scan {
		for j := i + 1; j < len(scan); j++ {
			mag := int(math.Round(Magnitude(getDirection(scan[j], beac))))
			//newDir := getNormalizedDirection(scan[j], beac)
			scanDirs = append(scanDirs, [3][3]int{beac, scan[j], [3]int{mag, mag, mag}})
		}
	}
	return scanDirs
}

func getAllDistances(scan [][3]int) []beaconPair {
	var allDists []beaconPair
	for i, beac := range scan {
		for j := i + 1; j < len(scan); j++ {
			mag := Magnitude(getDirection(scan[j], beac))
			if mag != Magnitude(getDirection(scan[j], beac)) {
				fmt.Println("ERROR")
			}
			//newDir := getNormalizedDirection(scan[j], beac)
			allDists = append(allDists, beaconPair{beac, scan[j], mag})
		}
	}
	return allDists
}

func getScannerDir(step1, step2 [3]int) [3]int {
	return [3]int{step1[0] - step2[0], step1[1] - step2[1], step1[2] - step2[2]}
}

func Magnitude(v [3]int) float64 {
	return math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2]))
}

// rotate90 represents a 90-degree rotation matrix along one axis.
var rotate90 = map[string][3][3]int{
	"X": {
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	},
	"Y": {
		{0, 0, 1},
		{0, 1, 0},
		{-1, 0, 0},
	},
	"Z": {
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	},
}

// identityMatrix returns a 3x3 identity matrix.
func identityMatrix() [3][3]int {
	return [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

// multiplyMatrices multiplies two 3x3 matrices.
func multiplyMatrices(a, b [3][3]int) [3][3]int {

	var result [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// applyRotation applies a 3x3 rotation matrix to a 3D vector.
func applyRotation(rotation [3][3]int, vector [3]int) [3]int {
	result := [3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i] += rotation[i][j] * vector[j]
		}
	}
	return result
}

// findRotation calculates the 3x3 rotation matrix to align v1 to v2.
func findRotation(v1, v2 [3]int) [3][3]int {
	rotations := [][3][3]int{
		identityMatrix(),
		rotate90["X"], multiplyMatrices(rotate90["X"], rotate90["X"]),
		multiplyMatrices(rotate90["X"], multiplyMatrices(rotate90["X"], rotate90["X"])),
		rotate90["Y"], multiplyMatrices(rotate90["Y"], rotate90["Y"]),
		multiplyMatrices(rotate90["Y"], multiplyMatrices(rotate90["Y"], rotate90["Y"])),
		rotate90["Z"], multiplyMatrices(rotate90["Z"], rotate90["Z"]),
		multiplyMatrices(rotate90["Z"], multiplyMatrices(rotate90["Z"], rotate90["Z"])),
	}

	for _, rz := range rotations {
		for _, ry := range rotations {
			for _, rx := range rotations {
				combined := multiplyMatrices(rz, multiplyMatrices(ry, rx))
				if applyRotation(combined, v1) == v2 {
					return combined
				}
			}
		}
	}
	return [3][3]int{} // No valid rotation found
}

func CrossProduct(u, v [3]float64) [3]float64 {
	return [3]float64{
		u[1]*v[2] - u[2]*v[1], // x-component
		u[2]*v[0] - u[0]*v[2], // y-component
		u[0]*v[1] - u[1]*v[0], // z-component
	}
}

func DotProduct(u, v [3]float64) float64 {
	return u[0]*v[0] + u[1]*v[1] + u[2]*v[2]
}

func findOffset(scan1 [][3]int, scan2 [][3]int) conversion {
	sameBeaconMap1 := make(map[[3]int]bool)
	sameBeaconMap2 := make(map[[3]int]bool)
	scan1Dirs := getAllDistances(scan1)
	scan2Dirs := getAllDistances(scan2)
	scannerOffset := make(map[conversion]int)
	for _, one := range scan1Dirs {
		for _, two := range scan2Dirs {
			if one.distance == two.distance {
				sameBeaconMap1[one.beacon1] = true
				sameBeaconMap1[one.beacon2] = true
				sameBeaconMap2[two.beacon1] = true
				sameBeaconMap2[two.beacon2] = true
				dirOne := getDirection(one.beacon2, one.beacon1)
				dirTwo := getDirection(two.beacon2, two.beacon1)

				rotMat := findRotation(dirTwo, dirOne)
				if rotMat == [3][3]int{} {
					continue
				}
				//rotVec := divVec(dirTwo, dirOne)
				leftDir := getDirection(applyRotation(rotMat, two.beacon1), one.beacon1)
				rightDir := getDirection(applyRotation(rotMat, two.beacon2), one.beacon2)

				if leftDir == rightDir {
					scannerOffset[conversion{leftDir, rotMat}] += 1
					//fmt.Println(getScannerDir(dir1[0], dir2[0]))
				}
			}
		}
	}
	maximum := 0
	var retConv conversion
	for idx, count := range scannerOffset {
		if count > maximum {
			maximum = count
			retConv = idx

		}
	}
	return retConv
}

func copySL(in [][2]int) [][2]int {
	var ret [][2]int
	for _, i2 := range in {
		ret = append(ret, i2)
	}
	return ret
}

func main() {
	file, err := os.Open("2021/day19/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var scannerList [][][3]int
	currentScanner := -1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "scanner") {
			scannerList = append(scannerList, make([][3]int, 0))
			currentScanner++
		} else {
			split := strings.Split(line, ",")
			newBeacon := [3]int{utils.ToInt(split[0]), utils.ToInt(split[1]), utils.ToInt(split[2])}
			scannerList[currentScanner] = append(scannerList[currentScanner], newBeacon)
		}
	}
	//fmt.Println(scannerList)
	//samePairMap := make(map[[3]int]map[[3]int]bool)
	//scannerOffset := findFirstTwo(scannerList[0], scannerList[1])
	//fmt.Println(scannerOffset)
	offSetMap := make(map[[2]int]conversion)
	//scannerOffset2 := findOffset(scannerList[0], scannerList[1])
	for i, _ := range scannerList {
		for j := 0; j < len(scannerList); j++ {
			scannerOffset := findOffset(scannerList[i], scannerList[j])
			if scannerOffset.offset != [3]int{0, 0, 0} {
				offSetMap[[2]int{i, j}] = scannerOffset
			}
		}
	}
	for i := 1; i < len(scannerList); i++ {
		fmt.Println(0, i, offSetMap[[2]int{0, i}])
	}
	for i := 1; i < len(scannerList); i++ {
		if _, ok := offSetMap[[2]int{0, i}]; !ok {
			var allPath [][][2]int
			for ints, _ := range offSetMap {
				if ints[0] == 0 {
					allPath = append(allPath, [][2]int{ints})
				}
			}
			var foundPath [][2]int
		outer:
			for {
				var newAllPath [][][2]int
				for _, curPath := range allPath {
					for ints, _ := range offSetMap {
						if curPath[len(curPath)-1][1] == ints[0] {
							tmpSlice := copySL(curPath)
							tmpSlice = append(tmpSlice, ints)
							if ints[1] == i {
								foundPath = tmpSlice
								break outer
							}
							newAllPath = append(newAllPath, tmpSlice)
						}
					}
				}
				allPath = newAllPath
			}
			summedOffset := offSetMap[foundPath[0]]
			for pathPos := 1; pathPos < len(foundPath); pathPos++ {
				summedOffset.offset = addVec(summedOffset.offset, applyRotation(summedOffset.rotMat, offSetMap[foundPath[pathPos]].offset))
				summedOffset.rotMat = multiplyMatrices(summedOffset.rotMat, offSetMap[foundPath[pathPos]].rotMat)
			}
			offSetMap[[2]int{0, i}] = summedOffset
		}
	}
	for i := 1; i < len(scannerList); i++ {
		fmt.Println(0, i, offSetMap[[2]int{0, i}])
	}
	uniqueBeacons := make(map[[3]int]bool)
	for i, list := range scannerList {
		if i == 0 {
			for _, beacon := range list {
				uniqueBeacons[beacon] = true
			}
		} else {
			for _, beacon := range list {
				shiftedBeacon := addVec(offSetMap[[2]int{0, i}].offset, applyRotation(offSetMap[[2]int{0, i}].rotMat, beacon))
				uniqueBeacons[shiftedBeacon] = true
			}
		}
	}
	fmt.Println(len(uniqueBeacons))

}
