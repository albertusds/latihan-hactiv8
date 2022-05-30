package helpers

import (
	"fmt"
	"sort"
)

func Fight(fighter []int, power []int, initialPower int) int {
	var testMap = make(map[int]int)

	//asign to map
	for i, tf := range fighter {
		testMap[tf] = power[i]
	}

	sort.Ints(fighter)

	fmt.Println(len(fighter))

	for _, data := range fighter {
		if initialPower >= data {
			initialPower += testMap[data]
		} else {
			break
		}
	}

	fmt.Println("final power :", initialPower)
	return initialPower
}
