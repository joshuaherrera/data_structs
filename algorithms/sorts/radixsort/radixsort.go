package radixsort

import "fmt"

//helper to get max integer in a slice
func maxInt(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}

//helper to count sort by digits 0-9
func countSort(arr []int, exp int) {
	output := make([]int, len(arr))
	digits := make([]int, 10)

	//count digits by exponent
	for i := 0; i < len(arr); i++ {
		idx := (arr[i] / exp) % 10
		//fmt.Println(idx)
		digits[idx]++
	}

	//arrange count[i] so it contains actual pos
	//of digit in output. used to put digits
	//in right order in next loop.
	for i := 1; i < 10; i++ {
		digits[i] += digits[i-1]
	}

	//build output array
	for i := len(arr) - 1; i >= 0; i-- {
		output[digits[(arr[i]/exp)%10]-1] = arr[i]
		digits[(arr[i]/exp)%10]--
	}

	copy(arr, output)

}

func RadixSort(arr []int) {
	fmt.Println("Start: ", arr)
	if len(arr) == 0 {
		return
	}

	max := maxInt(arr)

	//use count sort for each digit
	//use exponent as 10^i for each digits place
	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}
	fmt.Println("End: ", arr)
}
