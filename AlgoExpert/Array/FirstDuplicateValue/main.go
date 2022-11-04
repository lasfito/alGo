/* <p>
  Given an array of integers between <span>1</span> and <span>n</span>,
  inclusive, where <span>n</span> is the length of the array, write a function
  that returns the first integer that appears more than once (when the array is
  read from left to right).
</p>
<p>
  In other words, out of all the integers that might occur more than once in the
  input array, your function should return the one whose first duplicate value
  has the minimum index.
</p>
<p>
  If no integer appears more than once, your function should return
  <span>-1</span>.
</p>
<p>Note that you're allowed to mutate the input array.</p> */

package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 1, 5, 3, 3, 2, 4}
	fmt.Println(FirstDuplicateValue(arr))
}

func FirstDuplicateValue(array []int) int {

	for _, v := range array {
		absV := int(math.Abs(float64(v))) - 1
		if array[absV] < 0 {
			return int(math.Abs(float64(v)))
		} else {
			array[absV] = int(math.Abs(float64(array[absV]))) * -1
		}

	}
	return -1
}
