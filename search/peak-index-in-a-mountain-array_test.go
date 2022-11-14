package main

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Add test cases.
		{name: "test empty", args: args{[]int{}}, want: 0},
		{name: "test len:1 arr", args: args{[]int{10}}, want: 0},
		{name: "test len:2 arr by desc", args: args{[]int{10, 8}}, want: 0},
		{name: "test len:2 arr by asc", args: args{[]int{8, 10}}, want: 1},
		{name: "test len>2 arr by asc", args: args{[]int{8, 9, 10}}, want: 2},
		{name: "test len>2 arr by desc", args: args{[]int{10, 9, 8}}, want: 0},
		{name: "test len>2 peak index", args: args{[]int{0, 2, 1, 0}}, want: 1},
		{name: "test out range index", args: args{[]int{3, 5, 3, 2, 0}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArrayV1(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArrayV2(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArrayV3(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
