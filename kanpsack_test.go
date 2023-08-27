package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func Test_knapsack(t *testing.T) {
	for _, tc := range []struct {
		maxW        int
		weights     []int
		vals        []int
		expectedVal int
		expectedIds []int
	}{
		{maxW: 5, weights: []int{3, 2, 2}, vals: []int{3, 2, 4}, expectedVal: 7, expectedIds: []int{0, 2}},
		{maxW: 5, weights: []int{3, 2, 2}, vals: []int{3, 4, 4}, expectedVal: 8, expectedIds: []int{1, 2}},
		{maxW: 3, weights: []int{3, 2, 1}, vals: []int{1, 2, 3}, expectedVal: 5, expectedIds: []int{1, 2}},
		{maxW: 3, weights: []int{3, 2, 1}, vals: []int{3, 2, 1}, expectedVal: 3, expectedIds: []int{0}},
		{maxW: 3, weights: []int{3, 2, 1}, vals: []int{3, 2, 2}, expectedVal: 4, expectedIds: []int{1, 2}},
		{maxW: 3, weights: []int{1, 2, 3}, vals: []int{2, 2, 3}, expectedVal: 4, expectedIds: []int{0, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			val, ids := knapsack(tc.maxW, tc.weights, tc.vals)
			assert.Equal(t, tc.expectedVal, val)
			assert.Equal(t, tc.expectedIds, ids)
		})
	}
}
