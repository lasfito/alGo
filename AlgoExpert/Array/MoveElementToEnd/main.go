/*   You're given an array of integers and an integer. Write a function that moves
all instances of that integer in the array to the end of the array and returns
the array.

The function should perform this in place (i.e., it should mutate the input
array) and doesn't need to maintain the order of the other integers.
*/

package main

import "fmt"

func main() {

	arr := []int{2, 1, 2, 2, 2, 3, 4, 2}
	fmt.Println(MoveElementToEnd(arr, 2))
}

func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	left := 0
	right := len(array) - 1

	for left < right {

		if array[left] == toMove && array[right] != toMove {
			array[left], array[right] = array[right], array[left]
		}
		if array[left] != toMove {
			left += 1
		}
		if array[right] == toMove {
			right -= 1
		}
	}
	return array

}
