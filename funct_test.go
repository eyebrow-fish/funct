package funct

import (
	"reflect"
	"testing"
)

func TestFunctOf(t *testing.T) {
	type args struct {
		any interface{}
	}
	tests := []struct {
		name      string
		args      args
		want      Funct
		wantPanic bool
	}{
		{"slice", args{[]int{1, 2}}, Funct{[]int{1, 2}}, false},
		{"map", args{map[int]int{1: 2}}, Funct{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if tt.wantPanic {
					recover()
				}
			}()
			if got := FunctOf(tt.args.any); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FunctOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
