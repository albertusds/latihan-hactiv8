package main

import (
	"fmt"
	"sort"
)

func main() {
	var testMap = make(map[int]int)

	var testFighter = []int{5, 3, 9, 8}
	var testPower = []int{2, 2, 3, 1}
	var testInitPower = 3

	//asign to map
	for i, tf := range testFighter {
		testMap[tf] = testPower[i]
		fmt.Println("testMap tf: ", testMap[tf])
	}

	sort.Ints(testFighter)

	fmt.Println(len(testFighter))

	for _, data := range testFighter {
		if testInitPower >= data {
			testInitPower += testMap[data]
		} else {
			fmt.Println("ellse")
			break
		}
	}

	fmt.Println("final power :", testInitPower)
}
