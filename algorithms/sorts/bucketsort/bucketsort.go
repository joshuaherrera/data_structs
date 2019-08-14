package bucketsort

import (
	"math"
	"sort"
)

func Sort(arr []int, k int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	//make buckets of length k to store uniformy distributed ints
	buckets := make([][]int, k)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	max := maxInt(arr)

	//distribute each integer into a bucket
	for _, v := range arr {
		idx := int(math.Floor(float64(k*v/max))) - 1
		if idx < 0 {
			idx = 0
		}
		buckets[idx] = append(buckets[idx], v)
	}
	for _, bucket := range buckets {
		sort.Ints(bucket)
	}

	//make a new slice to return, and concatenate the sorted
	//slices, in order
	sorted := make([]int, 0)
	for _, sortedBucket := range buckets {
		sorted = append(sorted, sortedBucket...)
	}

	return sorted
}

func maxInt(arr []int) int {
	var max int = arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}
