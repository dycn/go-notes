package p680

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	type input struct {
		input1 string
	}
	type test struct {
		param input
		want  bool
	}

	tests := []test{
		{param: input{"abca"}, want: true},
		// {param: input{input1: []int{3, 2, 4}, input2: 5}, want: []int{0, 1}},
		// {param: input{input1: []int{0, 8, 7, 3, 3, 4, 2}, input2: 11}, want: []int{1, 3}},
		// {param: input{input1: []int{0, 1}, input2: 1}, want: []int{0, 1}},
	}

	for _, v := range tests {
		result := ValidPalindrome(v.param.input1)
		fmt.Println(result)
		// if !reflect.DeepEqual(v.want, result) {
		// 	t.Errorf("want: %v, result: %v", v.want, result)
		// }
	}
}
