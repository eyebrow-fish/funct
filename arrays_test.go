package funct

import (
	"reflect"
	"testing"
)

func TestIntArray_Filter(t *testing.T) {
	tests := []struct {
		name    string
		a       IntArray
		f       func(int) bool
		wantOut IntArray
	}{
		{name: "filter odd", a: []int{1, 2, 3}, f: func(i int) bool { return i%2 == 0 }, wantOut: []int{2}},
		{name: "filter even", a: []int{1, 2, 3}, f: func(i int) bool { return i%2 == 1 }, wantOut: []int{1, 3}},
		{name: "pass all", a: []int{1, 2, 3}, f: func(i int) bool { return true }, wantOut: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.a.Filter(tt.f); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Filter() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
