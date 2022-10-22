/*   Write a function that takes in a non-empty array of distinct integers and an
integer representing a target sum. If any two numbers in the input array sum
up to the target sum, the function should return them in an array, in any
order. If no two numbers sum up to the target sum, the function should return
an empty array.
Note that the target sum has to be obtained by summing two different integers
in the array; you can't add a single integer to itself in order to obtain the
target sum.
You can assume that there will be at most one pair of numbers summing up to
the target sum. */

package main

import "fmt"

func main() {
	var input = []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	fmt.Println(TwoNumberSum(input, target))
}

func TwoNumberSum(array []int, target int) []int {
	numMap := make(map[int]bool)
	var answer = []int{}

	for _, n := range array {
		numMap[n] = true
	}

	for _, n := range array {
		missingNum := target - n
		if (numMap[missingNum]) && (n != missingNum) {
			answer = append(answer, n, missingNum)
			break

		}
	}

	return answer
}
