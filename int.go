package funct

type IntArray []int

func (a IntArray) Filter(pred func(i int) bool) (out IntArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a IntArray) Foreach(f func(i int) int) {
	for _, v := range a {
		f(v)
	}
}

func (a IntArray) ForeachIndexed(f func(i int, v int) int) {
	for i, v := range a {
		f(i, v)
	}
}

func (a IntArray) Reduce(f func(a, b int) int) (out int) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a IntArray) Fold(start int, f func(a, b int) int) int {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Int8Array []int8

func (a Int8Array) Filter(pred func(i int8) bool) (out Int8Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Int8Array) Foreach(f func(i int8) int8) {
	for _, v := range a {
		f(v)
	}
}

func (a Int8Array) ForeachIndexed(f func(i int, v int8) int8) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Int8Array) Reduce(f func(a, b int8) int8) (out int8) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Int8Array) Fold(start int8, f func(a, b int8) int8) int8 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Int16Array []int16

func (a Int16Array) Filter(pred func(i int16) bool) (out Int16Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Int16Array) Foreach(f func(i int16) int16) {
	for _, v := range a {
		f(v)
	}
}

func (a Int16Array) ForeachIndexed(f func(i int, v int16) int16) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Int16Array) Reduce(f func(a, b int16) int16) (out int16) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Int16Array) Fold(start int16, f func(a, b int16) int16) int16 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Int32Array []int32

func (a Int32Array) Filter(pred func(i int32) bool) (out Int32Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Int32Array) Foreach(f func(i int32) int32) {
	for _, v := range a {
		f(v)
	}
}

func (a Int32Array) ForeachIndexed(f func(i int, v int32) int32) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Int32Array) Reduce(f func(a, b int32) int32) (out int32) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Int32Array) Fold(start int32, f func(a, b int32) int32) int32 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Int64Array []int64

func (a Int64Array) Filter(pred func(i int64) bool) (out Int64Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Int64Array) Foreach(f func(i int64) int64) {
	for _, v := range a {
		f(v)
	}
}

func (a Int64Array) ForeachIndexed(f func(i int, v int64) int64) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Int64Array) Reduce(f func(a, b int64) int64) (out int64) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Int64Array) Fold(start int64, f func(a, b int64) int64) int64 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

