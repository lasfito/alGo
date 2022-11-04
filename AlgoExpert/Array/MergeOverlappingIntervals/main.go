/*
<p>
  Write a function that takes in a non-empty array of arbitrary intervals,
  merges any overlapping intervals, and returns the new intervals in no
  particular order.
</p>
<p>
  Each interval <span>interval</span> is an array of two integers, with
  <span>interval[0]</span> as the start of the interval and
  <span>interval[1]</span> as the end of the interval.
</p>
<p>
  Note that back-to-back intervals aren't considered to be overlapping. For
  example, <span>[1, 5]</span> and <span>[6, 7]</span> aren't overlapping;
  however, <span>[1, 6]</span> and <span>[6, 7]</span> <i>are</i> indeed
  overlapping.
</p>
<p>
  Also note that the start of any particular interval will always be less than
  or equal to the end of that interval.
</p> */

package main

import (
	"fmt"
	"sort"
)

func main() {

	intervals := [][]int{{1, 2},
		{3, 5}, {6, 8}, {4, 7},
		{9, 10}}
	fmt.Println(MergeOverlappingIntervals(intervals))
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {

	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 1; i < len(intervals); {

		prevS := intervals[i-1][0]
		prevE := intervals[i-1][1]
		start := intervals[i][0]
		end := intervals[i][1]

		isOverlapping := (start <= prevE) && (end >= start) && (end >= prevS)
		if isOverlapping {
			var newStart int
			var newEnd int
			if start < prevS {
				newStart = start
			} else {
				newStart = prevS
			}

			if end > prevE {
				newEnd = end
			} else {
				newEnd = prevE
			}

			merged := [][]int{{newStart, newEnd}}
			prevArr := intervals[:i-1]
			var newArr [][]int
			if i == len(intervals)-1 {
				newArr = merged

			} else {

				newArr = append(merged, intervals[i+1:]...)
			}
			intervals = append(prevArr, newArr...)

		} else {
			i++
		}
	}
	return intervals
}
