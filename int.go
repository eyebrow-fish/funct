package funct

import "strconv"

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

func (a IntArray) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a IntArray) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a IntArray) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a IntArray) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a IntArray) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a IntArray) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a IntArray) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a IntArray) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a IntArray) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(v))
	}
	return
}

func (a IntArray) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a IntArray) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a IntArray) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a IntArray) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a IntArray) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
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

func (a Int8Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Int8Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Int8Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Int8Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Int8Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Int8Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Int8Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Int8Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Int8Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(int(v)))
	}
	return
}

func (a Int8Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Int8Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Int8Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Int8Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a Int8Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
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

func (a Int16Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Int16Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Int16Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Int16Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Int16Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Int16Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Int16Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Int16Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Int16Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(int(v)))
	}
	return
}

func (a Int16Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Int16Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Int16Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Int16Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a Int16Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
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

func (a Int32Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Int32Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Int32Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Int32Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Int32Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Int32Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Int32Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Int32Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, v)
	}
	return
}

func (a Int32Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(int(v)))
	}
	return
}

func (a Int32Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Int32Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Int32Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Int32Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a Int32Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
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

func (a Int64Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Int64Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Int64Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Int64Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Int64Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Int64Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Int64Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Int64Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Int64Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(int(v)))
	}
	return
}

func (a Int64Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Int64Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Int64Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Int64Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a Int64Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}
