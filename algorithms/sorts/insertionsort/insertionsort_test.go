package insertionsort

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
	}

	for _, c := range cases {
		Sort(c.in)
		if reflect.DeepEqual(c.in, c.want) == false {
			t.Errorf("Sorting error: got %v, wanted %v\n", c.in, c.want)
		}
	}

}
