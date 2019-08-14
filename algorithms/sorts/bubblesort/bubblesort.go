package bubblesort

//best case space complexity: O(n), avg & worst: O(n^2)
//space: 1
//stable
func Sort(arr []int) {
	var swapped bool
	for {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}
