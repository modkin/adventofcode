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

func Task1() {
	fmt.Println(getNextTenRecipes(894501))
}
