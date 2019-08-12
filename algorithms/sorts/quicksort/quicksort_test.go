package quicksort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{9, 88, 5, 1, 3, 0, 11, -11, -50, -999, 777, 86, 13, 7, 21}, []int{-999, -50, -11, 0, 1, 3, 5, 7, 9, 11, 13, 21, 86, 88, 777}},
	}

	for _, c := range cases {
		Sort(c.in, 0, len(c.in)-1)
		if reflect.DeepEqual(c.in, c.want) == false {
			t.Errorf("Sorting error: got %v, wanted %v\n", c.in, c.want)
		}
	}
}
