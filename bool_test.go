package funct

import (
	"reflect"
	"testing"
)

func TestBoolArray_Filter(t *testing.T) {
	type args struct {
		pred func(bool) bool
	}
	tests := []struct {
		name    string
		a       BoolArray
		args    args
		wantOut BoolArray
	}{
		{"all", BoolArray{false, true}, args{func(i bool) bool { return true }}, BoolArray{false, true}},
		{"true", BoolArray{false, true}, args{func(i bool) bool { return i }}, BoolArray{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.a.Filter(tt.args.pred); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Filter() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestBoolArray_Reduce(t *testing.T) {
	type args struct {
		f func(bool, bool) bool
	}
	tests := []struct {
		name    string
		a       BoolArray
		args    args
		wantOut bool
	}{
		{"has true", BoolArray{false, true, false}, args{func(a bool, b bool) bool { return a || b }}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.a.Reduce(tt.args.f); gotOut != tt.wantOut {
				t.Errorf("Reduce() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestBoolArray_Fold(t *testing.T) {
	type args struct {
		start bool
		f     func(bool, bool) bool
	}
	tests := []struct {
		name string
		a    BoolArray
		args args
		want bool
	}{
		{"only false", BoolArray{false}, args{false, func(a bool, b bool) bool { return a || b }}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Fold(tt.args.start, tt.args.f); got != tt.want {
				t.Errorf("Fold() = %v, want %v", got, tt.want)
			}
		})
	}
}
