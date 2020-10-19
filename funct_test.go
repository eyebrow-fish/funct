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
		name    string
		args    args
		want    *Funct
		wantErr bool
	}{
		{"slice", args{[]int{}}, &Funct{[]int{}}, false},
		{"map", args{map[int]int{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FunctOf(tt.args.any)
			if (err != nil) != tt.wantErr {
				t.Errorf("FunctOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FunctOf() got = %v, want %v", got, tt.want)
			}
		})
	}
}
