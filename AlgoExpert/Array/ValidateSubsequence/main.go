package main

import "fmt"

func main() {

	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq := []int{1, 6, -1, 10}
	fmt.Print(IsValidSubsequence(arr, seq))
}

func IsValidSubsequence(array []int, sequence []int) bool {
	i, j := 0, 0

	for i < len(array) && j < len(sequence) {
		if array[i] == sequence[j] {
			j++
		}
		i++
	}
	return j == len(sequence)
}
