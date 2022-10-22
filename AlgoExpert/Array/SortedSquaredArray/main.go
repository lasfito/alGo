package main

import (
	"fmt"
)

func main() {

	arr := []int{-3, -2, -1, 1, 2, 3, 4, 5, 6}
	fmt.Println(SortedSquaredArray(arr))

}

func SortedSquaredArray(array []int) []int {

	left, right := 0, len(array)-1
	sqrArray := []int(nil)
	fmt.Println(array)

	for len(sqrArray) != len(array) {
		sqrLeft := array[left] * array[left]
		sqrRight := array[right] * array[right]
		fmt.Println(array[left], array[right])
		if sqrLeft >= sqrRight {
			sqrArray = append([]int{sqrLeft}, sqrArray...)
			left += 1
		}
		if sqrRight > sqrLeft {
			sqrArray = append([]int{sqrRight}, sqrArray...)
			right -= 1
		}
	}

	return sqrArray
}
