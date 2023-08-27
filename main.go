package main

import "fmt"

func main() {
	val, ids := knapsack(5, []int{3, 2, 2}, []int{3, 2, 4})
	fmt.Printf("val: %d, ids: %v\n", val, ids)
}

func knapsack(maxW int, weights []int, vals []int) (int, []int) {
	n := len(weights)
	nLim := n + 1
	k := make([][]int, nLim)
	wLim := maxW + 1
	for i := 0; i < nLim; i++ {
		k[i] = make([]int, wLim)
	}

	a := make([][][]int, nLim)
	for i := 0; i < nLim; i++ {
		a[i] = make([][]int, wLim)
	}

	for i := 1; i < nLim; i++ {
		weight := weights[i-1]
		for w := 1; w < wLim; w++ {
			p := k[i-1][w]
			if w < weight {
				k[i][w] = p
				a[i][w] = a[i-1][w]
			} else {
				c := vals[i-1] + k[i-1][w-weight]
				if c > p {
					k[i][w] = c
					a[i][w] = append(a[i-1][w-weight], i-1)
				} else {
					k[i][w] = p
					a[i][w] = a[i-1][w]
				}
			}
		}
		fmt.Printf("i: %v, k: %v\n", i, k)
	}

	return k[n][maxW], a[n][maxW]
}
