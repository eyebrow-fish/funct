package funct

type BoolArray []bool

func (a BoolArray) Filter(pred func(i bool) bool) (out BoolArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a BoolArray) Foreach(f func(i bool) bool) {
	for _, v := range a {
		f(v)
	}
}

func (a BoolArray) ForeachIndexed(f func(i int, v bool) bool) {
	for i, v := range a {
		f(i, v)
	}
}

func (a BoolArray) Reduce(f func(a, b bool) bool) (out bool) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a BoolArray) Fold(start bool, f func(a, b bool) bool) bool {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}
