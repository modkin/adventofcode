package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type gate struct {
	in  [2]string
	out string
	op  string
}

func performOp(values map[string]int, g gate) bool {
	if _, ok := values[g.in[0]]; !ok {
		return true
	}
	if _, ok := values[g.in[1]]; !ok {
		return true
	}
	if g.op == "AND" {
		values[g.out] = values[g.in[0]] & values[g.in[1]]
	} else if g.op == "OR" {
		values[g.out] = values[g.in[0]] | values[g.in[1]]
	} else if g.op == "XOR" {
		values[g.out] = values[g.in[0]] ^ values[g.in[1]]
	}
	return false
}

func copyGates(gates []gate) (outGates []gate) {
	for _, g := range gates {
		outGates = append(outGates, gate{g.in, g.out, g.op})
	}
	return
}

func count(in *[8]int, num int) (out int) {
	for _, i2 := range in {
		if i2 == num {
			out++
		}
	}
	return out
}

func incSwapIdx(idx *[8]int, maxIdx int) {
	run := true
	for run {
		run = false
		idx[7]++
		for i := 7; i >= 0; i-- {
			if idx[i] >= maxIdx {
				idx[i] = 0
				idx[i-1]++
			}
		}
		for _, i2 := range idx {
			if count(idx, i2) > 1 {
				run = true
			}
		}
	}
}

func swapGates(idx *[8]int, gates []gate) {
	gates[idx[0]].out, gates[idx[1]].out = gates[idx[1]].out, gates[idx[0]].out
	gates[idx[2]].out, gates[idx[3]].out = gates[idx[3]].out, gates[idx[2]].out
	gates[idx[4]].out, gates[idx[5]].out = gates[idx[5]].out, gates[idx[4]].out
	gates[idx[6]].out, gates[idx[7]].out = gates[idx[7]].out, gates[idx[6]].out

}

func calcZ(values map[string]int, maxZ int) string {
	var out string
	for z := maxZ - 1; z >= 0; z-- {
		if z < 10 {
			out += strconv.Itoa(values["z0"+strconv.Itoa(z)])
		} else {
			out += strconv.Itoa(values["z"+strconv.Itoa(z)])
		}
	}
	return out
}

func getWrongIdx(expectedResult, result string) int {
	for i := range len(expectedResult) {
		if expectedResult[len(expectedResult)-i-1] != result[len(result)-i-1] {
			return len(result) - i - 1
		}
	}
	return -1
}

func swapOut(swap1 string, swap2 string, gates []gate) {
	var idx0, idx1 int
	for i, g := range gates {
		if g.out == swap1 {
			idx0 = i
		}
		if g.out == swap2 {
			idx1 = i
		}
	}
	gates[idx0].out, gates[idx1].out = gates[idx1].out, gates[idx0].out
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day24/input")

	values := make(map[string]int)
	gates := []gate{}

	stillValues := true
	maxZ := 0

	for _, line := range lines {
		if line == "" {
			stillValues = false
			continue
		}
		if stillValues {
			split := strings.Split(line, " ")
			values[strings.Trim(split[0], ":")] = utils.ToInt(split[1])

		} else {
			splitLine := strings.Split(line, " ")
			newGate := gate{[2]string{splitLine[0], splitLine[2]}, splitLine[4], splitLine[1]}
			gates = append(gates, newGate)
		}
	}

	rungates := func() {
		goOn := true
		for goOn {
			goOn = false
			for _, g := range gates {
				if performOp(values, g) {
					goOn = true
				}
			}
		}
	}
	rungates()

	for s, _ := range values {
		if string(s[0]) == "z" {
			maxZ++
		}
	}

	fmt.Println()
	var zOut string
	var xIn string
	var yIn string
	for z := maxZ - 1; z >= 0; z-- {
		if z < 10 {
			fmt.Print(values["z0"+strconv.Itoa(z)])
			zOut += strconv.Itoa(values["z0"+strconv.Itoa(z)])
			xIn += strconv.Itoa(values["x0"+strconv.Itoa(z)])
			yIn += strconv.Itoa(values["y0"+strconv.Itoa(z)])
		} else {
			fmt.Print(values["z"+strconv.Itoa(z)])
			zOut += strconv.Itoa(values["z"+strconv.Itoa(z)])
			if z < maxZ-2 {
				xIn += strconv.Itoa(values["x"+strconv.Itoa(z)])
				yIn += strconv.Itoa(values["y"+strconv.Itoa(z)])
			}
		}
	}
	part1, _ := strconv.ParseInt(zOut, 2, 64)
	fmt.Println()
	fmt.Println(part1)

	x, _ := strconv.ParseInt(xIn, 2, 64)
	y, _ := strconv.ParseInt(yIn, 2, 64)

	expectedResult := strconv.FormatInt(x+y, 2)
	fmt.Println("  " + xIn)
	fmt.Println("  " + yIn)
	fmt.Println(zOut)
	fmt.Println(" " + expectedResult)

	//wrongIdx := []int{}

	getDeps := func(output string) (bool, gate) {
		for _, g := range gates {
			if g.out == output {
				return true, g
			}
		}
		return false, gate{}
	}

	//origGates := copyGates(gates)
	//swapOut("scf", "z16")
	//swapOut("dtn", "prd")
	swapOut("z16", "hmk", gates)
	swapOut("z33", "fcd", gates)
	swapOut("z20", "fhp", gates)
	swapOut("rvf", "tpc", gates)
	test := []string{"z16", "hmk", "z33", "fcd", "z20", "fhp", "rvf", "tpc"}
	slices.Sort(test)
	fmt.Println(strings.Join(test, ","))
	rungates()
	zOut = calcZ(values, maxZ)

	wrongIdx := getWrongIdx(expectedResult, zOut)
	idxChar := "z" + fmt.Sprintf("%02d", wrongIdx)

	idxChar = "z45"
	depsSlice := []string{idxChar}
	for len(depsSlice) > 0 {
		fmt.Println(depsSlice)
		newDeepSlice := []string{}
		for _, g := range depsSlice {
			found, dep := getDeps(g)
			if found {
				newDeepSlice = append(newDeepSlice, dep.in[0])
				newDeepSlice = append(newDeepSlice, dep.op)
				newDeepSlice = append(newDeepSlice, dep.in[1])
			}
		}
		depsSlice = newDeepSlice
	}

	fmt.Println(zOut)
	fmt.Println(expectedResult)

	if zOut == expectedResult {
		fmt.Println("WUHU")
	} else {
		fmt.Println("WIRU")
	}
	//
	//fmt.Println("old")
	//newWrongIdx := getWrongIdx(expectedResult, zOut)
	//for newWrongIdx >= wrongIdx {
	//	for _, g := range gates {
	//		if g.out == deps[0] {
	//			continue
	//		} else {
	//			swapOut(deps[0], g.out)
	//			rungates()
	//			zOut = calcZ(values, maxZ)
	//			newWrongIdx = getWrongIdx(expectedResult, zOut)
	//			if newWrongIdx < wrongIdx {
	//				fmt.Println("maybe", newWrongIdx, deps[0], g.out)
	//			}
	//			swapOut(deps[0], g.out)
	//		}
	//	}
	//	fmt.Println("didn't work")
	//	break
	//}
	//fmt.Println(getDeps(deps[0]))
	//fmt.Println(getDeps(deps[1]))
	//
	//rungates()
	//
	//zOut = calcZ(values, maxZ)
	//fmt.Println(zOut)

	//for i, g := range gates {
	//	if g != origGates[i] {
	//		fmt.Println(g, origGates[i])
	//		fmt.Println("ERROR")
	//		break
	//	}
	//}

}
