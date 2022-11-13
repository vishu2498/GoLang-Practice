package main

import (
	"fmt"
	"sort"
)

/*
input: 6,6,6,5,5,3,1,1,1,1,1,1,1
output: 3,5,6,1
*/

func main() {
	arr := []int{6, 6, 6, 5, 5, 3, 1, 1, 1, 1, 1, 1, 1}

	tempMap := make(map[int]int)
	for _, value := range arr {
		if _, ok := tempMap[value]; ok {
			tempMap[value]++
		} else {
			tempMap[value] = 1
		}
	}

	fmt.Println(tempMap)

	var tempArr []int
	for number := range tempMap {
		tempArr = append(tempArr, number)
	}
	
	sort.Slice(tempArr, func(i, j int) bool {
		return tempMap[tempArr[i]] < tempMap[tempArr[j]]
	})

	fmt.Println(tempArr)
}
