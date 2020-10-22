package funct

type ByteArray []byte

func (a ByteArray) Filter(pred func(i byte) bool) (out ByteArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a ByteArray) Foreach(f func(i byte) byte) {
	for _, v := range a {
		f(v)
	}
}

func (a ByteArray) ForeachIndexed(f func(i int, v byte) byte) {
	for i, v := range a {
		f(i, v)
	}
}

func (a ByteArray) Reduce(f func(a, b byte) byte) (out byte) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a ByteArray) Fold(start byte, f func(a, b byte) byte) byte {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a ByteArray) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a ByteArray) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a ByteArray) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a ByteArray) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a ByteArray) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a ByteArray) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a ByteArray) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a ByteArray) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a ByteArray) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, string(v))
	}
	return
}

func (a ByteArray) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a ByteArray) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, v)
	}
	return
}

func (a ByteArray) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a ByteArray) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a ByteArray) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}
