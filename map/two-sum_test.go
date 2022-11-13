package main

import (
	"reflect"
	"testing"
)

func TestTwoSumWithIndex(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		// TODO: Add test cases.
		{
			"test ok",
			args{
				[]int{2, 2, 2, -2}, 0,
			},
			[][2]int{{0, 3}, {1, 3}, {2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumWithIndex(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumWithIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
