package main

import (
	"fmt"
	"log"
)

func main() {

	//TODO: recive input first line
	var n, k int
	fmt.Println("Input:")
	fmt.Scan(&n, &k)
	if n <= 0 || k <= 0 {
		log.Fatal("Error Input")
	}

	//TODO: recive input second line
	positions := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&positions[i])
	}

	result := maxChickensProtected(n, k, positions)
	fmt.Printf("Output %d\n", result)
}

func maxChickensProtected(n, k int, positions []int) int {
	maxChickens := 0
	start := 0

	//TODO: first loop all position
	for end := 0; end < n; end++ {
		//TODO: loop for position <= k
		for positions[end]-positions[start] >= k {
			start++
		}

		//fmt.Printf("start %d , end %d\n", positions[start], positions[end])

		//TODO: check distance between chickens
		chickens := end - start + 1
		//fmt.Println("total:", chickens)

		//TODO: find most chockens rescue
		if chickens > maxChickens {
			maxChickens = chickens
		}
	}

	return maxChickens
}
