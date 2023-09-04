package main

import "fmt"

func main() {
	val, ids := knapsack(5, []int{3, 2, 2}, []int{3, 2, 4})
	fmt.Printf("val: %d, ids: %v\n", val, ids)
}

func knapsack(maxW int, weights []int, vals []int) (int, []int) {
	n := len(weights)
	wLim := maxW + 1
	var k [2][]int
	for i := 0; i < 2; i++ {
		k[i] = make([]int, wLim)
	}

	var a [2][][]int
	for i := 0; i < 2; i++ {
		a[i] = make([][]int, wLim)
	}

	for i := 0; i < n; i++ {
		weight := weights[i]
		for w := weight; w < wLim; w++ {
			wDiff := w - weight
			p := k[0][w]
			c := vals[i] + k[0][wDiff]
			if c > p {
				k[1][w] = c
				a[1][w] = append(a[0][wDiff], i)
			}
		}
		copy(k[0], k[1])
		copy(a[0], a[1])
	}

	return k[1][maxW], a[1][maxW]
}
