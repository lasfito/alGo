/*   Write a function that takes in a non-empty array of distinct integers and an
integer representing a target sum. The function should find all triplets in
the array that sum up to the target sum and return a two-dimensional array of
all these triplets. The numbers in each triplet should be ordered in ascending
order, and the triplets themselves should be ordered in ascending order with
respect to the numbers they hold. */

package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{1, 2, 3}
	fmt.Println(ThreeNumberSum(arr, 7))
}

func ThreeNumberSum(array []int, target int) [][]int {

	twoSums := make(map[int][][]int)
	sort.Ints(array)
	// indexMitad := math.Ceil(float64(len(array)) / 2)

	for i, v := range array {
		// if i > int(indexMitad) {
		// 	continue
		// }
		for j, v2 := range array {
			if i == j {
				continue
			}
			twoSums[v+v2] = append(twoSums[v+v2], []int{v, v2})
		}
	}

	var triplets [][]int

	for _, v := range array {
		needVal := twoSums[target-v]
		usedNumbers := make(map[int]bool)

		for _, arr := range needVal {

			prime := arr[0]
			secuns := arr[1]

			if (prime == v) || (secuns == v) {
				continue
			}
			if (usedNumbers[prime] != false) && (usedNumbers[secuns] != false) {
				currTriplet := []int{v, prime, secuns}
				sort.Ints(currTriplet)
				triplets = append(triplets, currTriplet)
			}
			usedNumbers[prime] = true
			usedNumbers[secuns] = true

		}

	}

	if len(triplets) == 0 {
		return triplets
	}

	for _, arr := range triplets {
		sort.Ints(arr)
	}

	for i, arr := range triplets {
		if i == len(triplets)-1 {
			continue
		}
		if arr[0] > triplets[i+1][0] {
			triplets[i], triplets[i+1] = triplets[i+1], triplets[i]
		}
	}

	for i, arr := range triplets {
		if i == len(triplets)-1 {
			continue
		}
		if arr[1] > triplets[i+1][1] && arr[0] == triplets[i+1][0] {
			triplets[i], triplets[i+1] = triplets[i+1], triplets[i]
		}
	}

	return triplets
}
