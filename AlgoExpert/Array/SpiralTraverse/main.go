/* <p>
  Write a function that takes in an n x m two-dimensional array (that can be
  square-shaped when n == m) and returns a one-dimensional array of all the
  array's elements in spiral order.
</p>

<p>
  Spiral order starts at the top left corner of the two-dimensional array, goes
  to the right, and proceeds in a spiral pattern all the way until every element
  has been visited.
</p> */

package main

import (
	"fmt"
)

func main() {

	arr := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7}}

	fmt.Println("final", SpiralTraverse(arr))

}

func SpiralTraverse(array [][]int) []int {
	// Write your code here.

	totalRows := len(array)
	totalCols := len(array[0])
	scannedRight := 0
	scannedDown := 0
	scannedLeft := 0
	scannedUp := 0
	ans := []int{}

	for len(ans) != totalCols*totalRows {

		if scannedDown == scannedUp && scannedRight == scannedLeft {
			for i := scannedUp; i < totalCols-scannedDown; i++ {
				ans = append(ans, array[scannedRight][i])
			}

			scannedRight++
		}

		if scannedRight > scannedLeft && scannedUp == scannedDown {

			for i := scannedRight; i < totalRows-scannedLeft; i++ {
				ans = append(ans, array[i][totalCols-scannedDown-1])
			}

			scannedDown++

		}

		if scannedLeft < scannedRight && scannedDown > scannedUp {
			for i := totalCols - scannedDown; i > scannedUp; i-- {
				ans = append(ans, array[totalRows-scannedLeft-1][i-1])
			}

			scannedLeft++
		}

		if scannedLeft == scannedRight && scannedUp < scannedDown {

			for i := totalRows - scannedLeft; i > scannedRight; i-- {
				fmt.Println(array[i-1][scannedUp])
				ans = append(ans, array[i-1][scannedUp])
			}

			scannedUp++
		}

	}

	return ans
}
