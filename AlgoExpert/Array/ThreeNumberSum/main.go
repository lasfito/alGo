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

	arr := []int{-8, -6, 12, 3, 1, 2, 5, 6}
	fmt.Println(ThreeNumberSum(arr, 0))
}

func ThreeNumberSum(array []int, target int) [][]int {

	sort.Ints(array)
	triplets := [][]int{}

	for i, v := range array {

		if i == len(array)-1 {
			break
		}
		left := i + 1
		right := len(array) - 1

		for left < right {
			currentSum := v + array[left] + array[right]
			if currentSum == target {
				currentTriplet := []int{v, array[left], array[right]}
				triplets = append(triplets, currentTriplet)
				left += 1
				right -= 1
			}
			if currentSum < target {
				left += 1
			}
			if currentSum > target {
				right -= 1
			}

		}

	}

	return triplets
}
