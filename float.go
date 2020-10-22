package funct

import (
	"fmt"
	"strconv"
)

type Float32Array []float32

func (a Float32Array) Filter(pred func(float32) bool) (out Float32Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Float32Array) Foreach(f func(float32) float32) {
	for _, v := range a {
		f(v)
	}
}

func (a Float32Array) ForeachIndexed(f func(int, float32) float32) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Float32Array) Reduce(f func(float32, float32) float32) (out float32) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Float32Array) Fold(start float32, f func(float32, float32) float32) float32 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a Float32Array) Find(pred func(float32) bool) (float32, error) {
	for _, v := range a {
		if pred(v) {
			return v, nil
		}
	}
	return 0, fmt.Errorf("no value found")
}

func (a Float32Array) AnyMatch(pred func(float32) bool) bool {
	for _, v := range a {
		if pred(v) {
			return true
		}
	}
	return false
}

func (a Float32Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Float32Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Float32Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Float32Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Float32Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Float32Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Float32Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Float32Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Float32Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.FormatFloat(float64(v), 'f', 6, 32))
	}
	return
}

func (a Float32Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Float32Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Float32Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Float32Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a Float32Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}

type Float64Array []float64

func (a Float64Array) Filter(pred func(float64) bool) (out Float64Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Float64Array) Foreach(f func(float64) float64) {
	for _, v := range a {
		f(v)
	}
}

func (a Float64Array) ForeachIndexed(f func(int, float64) float64) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Float64Array) Reduce(f func(float64, float64) float64) (out float64) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Float64Array) Fold(start float64, f func(float64, float64) float64) float64 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a Float64Array) Find(pred func(float64) bool) (float64, error) {
	for _, v := range a {
		if pred(v) {
			return v, nil
		}
	}
	return 0, fmt.Errorf("no value found")
}

func (a Float64Array) AnyMatch(pred func(float64) bool) bool {
	for _, v := range a {
		if pred(v) {
			return true
		}
	}
	return false
}

func (a Float64Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Float64Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Float64Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Float64Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Float64Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Float64Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Float64Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Float64Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Float64Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.FormatFloat(v, 'f', 6, 64))
	}
	return
}

func (a Float64Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Float64Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Float64Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Float64Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a Float64Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}
