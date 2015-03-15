package qsort

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(vals []int, exp []int, e *testing.T) {
	Sort(vals)

	fmt.Println("result is:", vals)

	if !reflect.DeepEqual(vals, exp) {
		e.Error("\nresult is:", vals, "\nexpect is:", exp)
	}
}

func TestQSort0(e *testing.T) {
	assert ([]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5}, e)
}

func TestQSort1(e *testing.T) {
	assert ([]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5}, e)
}

func TestQSort2(e *testing.T) {
	assert ([]int{5, 4, 5, 2, 1},
			[]int{1, 2, 4, 5, 5}, e)
}


func TestQSort3(e *testing.T) {
	assert ([]int{5, 4},
			[]int{4, 5}, e)
}

func TestQSort4(e *testing.T) {
	assert ([]int{5},
			[]int{5}, e)
}

func TestQSort5(e *testing.T) {
	assert ([]int{},
			[]int{}, e)
}

