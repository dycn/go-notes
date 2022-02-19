package p1

import (
	"reflect"
	"testing"
)
func TestTwoSum(t *testing.T){
	type input struct {
		input1 []int
		input2 int
	}
	type test struct {
		param input
		want []int
	}

	tests := []test{
		{param: input{input1: []int{3, 2, 4}, input2: 6}, want: []int{1, 2}},
		{param: input{input1: []int{3, 2, 4}, input2: 5}, want: []int{0, 1}},
		{param: input{input1: []int{0, 8, 7, 3, 3, 4, 2}, input2: 11}, want: []int{1, 3}},
		{param: input{input1: []int{0, 1}, input2: 1}, want: []int{0, 1}},
	}

	for _, v := range tests {
		result := twoSum(v.param.input1, v.param.input2)
		if !reflect.DeepEqual(v.want, result) {
			t.Errorf("want: %v, result: %v", v.want, result)
		}
	}
}