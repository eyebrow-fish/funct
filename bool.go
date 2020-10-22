package funct

import "fmt"

type BoolArray []bool

func (a BoolArray) Filter(pred func(bool) bool) (out BoolArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a BoolArray) Foreach(f func(bool) bool) {
	for _, v := range a {
		f(v)
	}
}

func (a BoolArray) ForeachIndexed(f func(int, bool) bool) {
	for i, v := range a {
		f(i, v)
	}
}

func (a BoolArray) Reduce(f func(bool, bool) bool) (out bool) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a BoolArray) Fold(start bool, f func(bool, bool) bool) bool {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a BoolArray) Find(pred func(bool) bool) (bool, error) {
	for _, v := range a {
		if pred(v) {
			return v, nil
		}
	}
	return false, fmt.Errorf("no value found")
}

func (a BoolArray) AnyMatch(pred func(bool) bool) bool {
	for _, v := range a {
		if pred(v) {
			return true
		}
	}
	return false
}
