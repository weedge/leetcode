package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantPos int
	}{
		// TODO: Add test cases.
		{
			name:    "ok",
			args:    args{arr: []int{1, 3}, target: 3},
			wantPos: 1,
		},
		{
			name:    "ok",
			args:    args{arr: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			wantPos: 4,
		},
		{
			name:    "no",
			args:    args{arr: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
			wantPos: -1,
		},
		{
			name:    "no",
			args:    args{arr: []int{2}, target: 0},
			wantPos: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPos := search(tt.args.arr, tt.args.target); gotPos != tt.wantPos {
				t.Errorf("search() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}
