package funct

import (
	"fmt"
	"reflect"
)

// Funct is a functional interface.
// Funct wraps an existing slice and allows specific transformations
// and collections that are typical to the functional programming paradigm.
//
// The below creates a slice of even ints
//
//	f, _ := FunctOf([]int{1, 2, 3, 4})
//	f.Filter(func(i int) bool) { return i % 2 }).Collect()
type Funct struct {
	data interface{}
}

// FunctOf is the primary constructor of Funct.
// Allows a slice of any kind.
//
//	FuncOf([]string{"hello", "world"})
//
// Returns a pointer of the created Funct or nil if it could not be
// created. In the case of failure, such as a bad parameter, an error
// will be provided.
func FunctOf(any interface{}) (*Funct, error) {
	t := reflect.TypeOf(any)
	if t.Kind() != reflect.Slice {
		return nil, fmt.Errorf("could not create error from {%v}", any)
	}
	return &Funct{data: any}, nil
}
