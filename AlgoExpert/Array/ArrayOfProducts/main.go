/* <p>
  Write a function that takes in a non-empty array of integers and returns an
  array of the same length, where each element in the output array is equal to
  the product of every other number in the input array.
</p>
<p>
  In other words, the value at <span>output[i]</span> is equal to the product of
  every number in the input array other than <span>input[i]</span>.
</p>
<p>Note that you're expected to solve this problem without using division.</p> */

package main

import "fmt"

func main() {

	arr := []int{5, 1, 4, 2}
	fmt.Println(ArrayOfProducts(arr))
}

func ArrayOfProducts(array []int) []int {
	var output []int
	acc := 1
	// izq -> der
	for i := 0; i < len(array); i++ {
		output = append(output, acc)
		acc *= array[i]
	}
	acc = 1
	//der -> izq
	for i := len(array) - 1; i >= 0; i-- {
		output[i] *= acc
		acc *= array[i]
	}
	return output
}
