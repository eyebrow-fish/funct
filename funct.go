package funct

import (
	"log"
	"reflect"
)

// Funct is a functional interface.
// Funct wraps an existing slice and allows specific transformations
// and collections that are typical to the functional programming paradigm.
type Funct struct {
	data interface{}
}

// FunctOf is the primary constructor of Funct.
// Allows a slice of any kind.
//
//	FuncOf([]string{"hello", "world"})
//
// Returns a Funct of the data provided, unless the parameter is not a valid type.
func FunctOf(any interface{}) Funct {
	t := reflect.TypeOf(any)
	if t.Kind() != reflect.Slice {
		log.Panicf("could not create error from {%v}", any)
	}
	return Funct{data: any}
}

// Filter is a function of Funct for filtering out elements in our data
// which do not meet the predicate function provided.
//
// In the example below the Filter only accepts even ints.
//
//	f.Filter(func(i int) bool { return i % 2 == 0 })
//
// Returns a new Funct containing the elements passed the predicate.
func (f Funct) Filter(handler interface{}) Funct {
	h := reflect.ValueOf(handler)
	if h.Kind() != reflect.Func {
		log.Panicf("argument must be a function, but was {%v}", handler)
	}
	d := reflect.ValueOf(f.data)
	var data []interface{}
	for i := 0; i < d.Len(); i++ {
		v := d.Index(i)
		if h.Call([]reflect.Value{v})[0].Bool() {
			data = append(data, v.Interface())
		}
	}
	return Funct{data}
}
