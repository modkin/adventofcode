package day14

import "fmt"

func getNextRecipes(recs []int, elv1 int, elv2 int) (newIdx1 int, newIdx2 int, newRecs []int) {
	sum := recs[elv1] + recs[elv2]
	if sum > 9 {
		firstNew := int(sum / 10)
		secondNew := sum % 10
		newRecs = append(recs, firstNew, secondNew)
	} else {
		newRecs = append(recs, sum)
	}
	newIdx1 = (elv1 + recs[elv1] + 1) % len(newRecs)
	newIdx2 = (elv2 + recs[elv2] + 1) % len(newRecs)
	return
}

func getNextTenRecipes(tries int) (ret []int) {
	elv1 := 0
	elv2 := 1
	recipies := []int{3, 7}
	for len(recipies) < tries+10 {
		elv1, elv2, recipies = getNextRecipes(recipies, elv1, elv2)
	}
	ret = append(ret, recipies[(tries):tries+10]...)
	return
}

func searchRec(target []int) int {
	elv1 := 0
	elv2 := 1
	recipies := []int{3, 7}
	count := 0
	goOn := true
	for goOn {
		elv1, elv2, recipies = getNextRecipes(recipies, elv1, elv2)
		if len(recipies) < len(target)+1 {
			goOn = true
		} else {
			count++
			goOn = NotEqual(target, recipies[(count):count+len(target)])
		}
	}

	return count
}

func NotEqual(src []int, dst []int) bool {
	if len(dst) < len(src) {
		return true
	}
	for idx, _ := range src {
		if src[idx] != dst[idx] {
			return true
		}
	}
	return false
}

func search(target []int) (times int) {
	ret := make([]int, 10)
	i := 10
	for NotEqual(target, ret) {
		ret = getNextTenRecipes(i)
		i++
	}
	return (i - 1)
}

func Task1() {
	fmt.Println(getNextTenRecipes(894501))
}

func Task2() {
	//target1 := []int{5, 1, 5, 8, 9}
	target2 := []int{8, 9, 4, 5, 0, 1}
	fmt.Println(searchRec(target2))
}
