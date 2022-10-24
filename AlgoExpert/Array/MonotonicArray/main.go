/*   Write a function that takes in an array of integers and returns a boolean
representing whether the array is monotonic.
An array is said to be monotonic if its elements, from left to right, are
entirely non-increasing or entirely non-decreasing.

Non-increasing elements aren't necessarily exclusively decreasing; they simply
don't increase. Similarly, non-decreasing elements aren't necessarily
exclusively increasing; they simply don't decrease.
*/

package main

import "fmt"

func main() {
	arr := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}

	fmt.Println(IsMonotonic(arr))
}
func IsMonotonic(array []int) bool {
	// Write your code here.
	isMono := true
	isIncreasing := false
	isDecresing := false

	for i, v := range array {
		if i == len(array)-1 || isMono == false {
			break
		}
		if v < array[i+1] && isIncreasing == false {
			isDecresing = true
		}
		if v > array[i+1] && isDecresing == false {
			isIncreasing = true
		}
		if isDecresing && v > array[i+1] {
			isMono = false
		}
		if isIncreasing && v < array[i+1] {
			isMono = false
		}

	}
	return isMono
}
