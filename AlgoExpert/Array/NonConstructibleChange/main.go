/*
  Given an array of positive integers representing the values of coins in your
  possession, write a function that returns the minimum amount of change (the
  minimum sum of money) that you <b>cannot</b> create. The given coins can have
  any positive integer value and aren't necessarily unique (i.e., you can have
  multiple coins of the same value).
*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	coins := []int{1, 1, 1, 1, 1}

	missing := NonConstructibleChange(coins)
	fmt.Println(missing)
}

func NonConstructibleChange(coins []int) int {

	if len(coins) == 0 {
		return 1
	}
	sort.Ints(coins)
	fmt.Println(coins)
	maxChange := 0
	if coins[0] != 1 {
		return 1
	}
	for _, v := range coins {
		if v > maxChange+1 {
			return maxChange + 1
		}
		maxChange += v

	}
	return maxChange + 1
}
