package mergersort

import (
	"fmt"
)

//best, avg, worst case space complexity: O(nlgn)
//space complexity : O(n)
//stable
func Sort(input []int, buffer []int) {
	fmt.Println(input)
	copy(buffer, input)
	splitMerge(buffer, 0, len(input), input)
}

func splitMerge(buffer []int, start int, end int, input []int) {
	if end-start < 2 {
		// run of size 1; sorted
		return
	}
	mid := (end + start) / 2

	splitMerge(input, start, mid, buffer)
	splitMerge(input, mid, end, buffer)
	merge(buffer, start, mid, end, input)
}

func merge(input []int, start int, mid int, end int, buffer []int) {
	i := start
	j := mid

	for k := start; k < end; k++ {
		if i < mid && (j >= end || input[i] <= input[j]) {
			buffer[k] = input[i]
			i++
		} else {
			buffer[k] = input[j]
			j++
		}
	}
	fmt.Println(buffer)
}
