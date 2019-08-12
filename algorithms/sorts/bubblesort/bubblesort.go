package bubblesort

//for ease of use, compatible with integers only
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
