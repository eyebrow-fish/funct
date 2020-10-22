package funct

type Float32Array []float32

func (a Float32Array) Filter(pred func(i float32) bool) (out Float32Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Float32Array) Foreach(f func(i float32) float32) {
	for _, v := range a {
		f(v)
	}
}

func (a Float32Array) ForeachIndexed(f func(i int, v float32) float32) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Float32Array) Reduce(f func(a, b float32) float32) (out float32) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Float32Array) Fold(start float32, f func(a, b float32) float32) float32 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Float64Array []float64

func (a Float64Array) Filter(pred func(i float64) bool) (out Float64Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Float64Array) Foreach(f func(i float64) float64) {
	for _, v := range a {
		f(v)
	}
}

func (a Float64Array) ForeachIndexed(f func(i int, v float64) float64) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Float64Array) Reduce(f func(a, b float64) float64) (out float64) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Float64Array) Fold(start float64, f func(a, b float64) float64) float64 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}
