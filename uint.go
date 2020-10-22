package funct

type UintArray []uint

func (a UintArray) Filter(pred func(i uint) bool) (out UintArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a UintArray) Foreach(f func(i uint) uint) {
	for _, v := range a {
		f(v)
	}
}

func (a UintArray) ForeachIndexed(f func(i int, v uint) uint) {
	for i, v := range a {
		f(i, v)
	}
}

func (a UintArray) Reduce(f func(a, b uint) uint) (out uint) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a UintArray) Fold(start uint, f func(a, b uint) uint) uint {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Uint8Array []uint8

func (a Uint8Array) Filter(pred func(i uint8) bool) (out Uint8Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Uint8Array) Foreach(f func(i uint8) uint8) {
	for _, v := range a {
		f(v)
	}
}

func (a Uint8Array) ForeachIndexed(f func(i int, v uint8) uint8) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Uint8Array) Reduce(f func(a, b uint8) uint8) (out uint8) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Uint8Array) Fold(start uint8, f func(a, b uint8) uint8) uint8 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Uint16Array []uint16

func (a Uint16Array) Filter(pred func(i uint16) bool) (out Uint16Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Uint16Array) Foreach(f func(i uint16) uint16) {
	for _, v := range a {
		f(v)
	}
}

func (a Uint16Array) ForeachIndexed(f func(i int, v uint16) uint16) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Uint16Array) Reduce(f func(a, b uint16) uint16) (out uint16) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Uint16Array) Fold(start uint16, f func(a, b uint16) uint16) uint16 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Uint32Array []uint32

func (a Uint32Array) Filter(pred func(i uint32) bool) (out Uint32Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Uint32Array) Foreach(f func(i uint32) uint32) {
	for _, v := range a {
		f(v)
	}
}

func (a Uint32Array) ForeachIndexed(f func(i int, v uint32) uint32) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Uint32Array) Reduce(f func(a, b uint32) uint32) (out uint32) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Uint32Array) Fold(start uint32, f func(a, b uint32) uint32) uint32 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

type Uint64Array []uint64

func (a Uint64Array) Filter(pred func(i uint64) bool) (out Uint64Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Uint64Array) Foreach(f func(i uint64) uint64) {
	for _, v := range a {
		f(v)
	}
}

func (a Uint64Array) ForeachIndexed(f func(i int, v uint64) uint64) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Uint64Array) Reduce(f func(a, b uint64) uint64) (out uint64) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Uint64Array) Fold(start uint64, f func(a, b uint64) uint64) uint64 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}
