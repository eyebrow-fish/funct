package funct

type RuneArray []rune

func (a RuneArray) Filter(pred func(i rune) bool) (out RuneArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a RuneArray) Foreach(f func(i rune) rune) {
	for _, v := range a {
		f(v)
	}
}

func (a RuneArray) ForeachIndexed(f func(i int, v rune) rune) {
	for i, v := range a {
		f(i, v)
	}
}

func (a RuneArray) Reduce(f func(a, b rune) rune) (out rune) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a RuneArray) Fold(start rune, f func(a, b rune) rune) rune {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a RuneArray) Join() (out string) {
	for _, v := range a {
		out += string(v)
	}
	return
}

func (a RuneArray) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a RuneArray) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a RuneArray) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a RuneArray) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a RuneArray) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a RuneArray) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a RuneArray) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, v)
	}
	return
}

func (a RuneArray) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a RuneArray) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, string(v))
	}
	return
}

func (a RuneArray) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a RuneArray) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a RuneArray) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a RuneArray) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a RuneArray) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}
