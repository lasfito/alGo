/*   Write a function that takes in two non-empty arrays of integers, finds the
pair of numbers (one from each array) whose absolute difference is closest to
zero, and returns an array containing these two numbers, with the number from
the first array in the first position.
Note that the absolute difference of two integers is the distance between
them on the real number line. For example, the absolute difference of -5 and 5
is 10, and the absolute difference of -5 and -4 is 1.

You can assume that there will only be one pair of numbers with the smallest
difference.
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	arr1 := []int{-1, 3, 5, 10, 20, 28}
	arr2 := []int{15, 17, 26, 134, 135}
	fmt.Println(SmallestDifference(arr1, arr2))
}

func SmallestDifference(arr1, arr2 []int) []int {

	sort.Ints(arr1)
	sort.Ints(arr2)
	var currDiff int = arr1[len(arr1)-1] + arr2[len(arr2)-1]
	var currAns []int
	i, j := 0, 0

	for i != len(arr1) && j != len(arr2) {
		up := arr1[i]
		down := arr2[j]
		curr := int(math.Abs(float64(up - down)))
		if curr == 0 {
			currAns = []int{up, down}
			break
		}
		if curr < currDiff {
			currDiff = curr
			currAns = []int{up, down}
		}
		if up < down {
			i += 1
		} else {
			j += 1
		}

	}

	return currAns
}
