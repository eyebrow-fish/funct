package funct

import "strconv"

type UintArray []uint

func (a UintArray) Filter(pred func(uint) bool) (out UintArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a UintArray) Foreach(f func(uint) uint) {
	for _, v := range a {
		f(v)
	}
}

func (a UintArray) ForeachIndexed(f func(int, uint) uint) {
	for i, v := range a {
		f(i, v)
	}
}

func (a UintArray) Reduce(f func(uint, uint) uint) (out uint) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a UintArray) Fold(start uint, f func(uint, uint) uint) uint {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a UintArray) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a UintArray) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a UintArray) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a UintArray) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a UintArray) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a UintArray) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a UintArray) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a UintArray) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a UintArray) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a UintArray) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(int(v)))
	}
	return
}

func (a UintArray) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a UintArray) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a UintArray) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a UintArray) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}

type Uint8Array []uint8

func (a Uint8Array) Filter(pred func(uint8) bool) (out Uint8Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Uint8Array) Foreach(f func(uint8) uint8) {
	for _, v := range a {
		f(v)
	}
}

func (a Uint8Array) ForeachIndexed(f func(int, uint8) uint8) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Uint8Array) Reduce(f func(uint8, uint8) uint8) (out uint8) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Uint8Array) Fold(start uint8, f func(uint8, uint8) uint8) uint8 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a Uint8Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, v)
	}
	return
}

func (a Uint8Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Uint8Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Uint8Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Uint8Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Uint8Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Uint8Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Uint8Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Uint8Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Uint8Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, string(v))
	}
	return
}

func (a Uint8Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Uint8Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Uint8Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a Uint8Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}

type Uint16Array []uint16

func (a Uint16Array) Filter(pred func(uint16) bool) (out Uint16Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Uint16Array) Foreach(f func(uint16) uint16) {
	for _, v := range a {
		f(v)
	}
}

func (a Uint16Array) ForeachIndexed(f func(int, uint16) uint16) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Uint16Array) Reduce(f func(uint16, uint16) uint16) (out uint16) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Uint16Array) Fold(start uint16, f func(uint16, uint16) uint16) uint16 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a Uint16Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Uint16Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Uint16Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Uint16Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Uint16Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Uint16Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Uint16Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Uint16Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Uint16Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Uint16Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(int(v)))
	}
	return
}

func (a Uint16Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Uint16Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Uint16Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}

func (a Uint16Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}

type Uint32Array []uint32

func (a Uint32Array) Filter(pred func(uint32) bool) (out Uint32Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Uint32Array) Foreach(f func(uint32) uint32) {
	for _, v := range a {
		f(v)
	}
}

func (a Uint32Array) ForeachIndexed(f func(int, uint32) uint32) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Uint32Array) Reduce(f func(uint32, uint32) uint32) (out uint32) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Uint32Array) Fold(start uint32, f func(uint32, uint32) uint32) uint32 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a Uint32Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Uint32Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Uint32Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Uint32Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Uint32Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Uint32Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Uint32Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Uint32Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Uint32Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Uint32Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(int(v)))
	}
	return
}

func (a Uint32Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Uint32Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Uint32Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Uint32Array) ToUint64Array() (out Uint64Array) {
	for _, v := range a {
		out = append(out, uint64(v))
	}
	return
}

type Uint64Array []uint64

func (a Uint64Array) Filter(pred func(uint64) bool) (out Uint64Array) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a Uint64Array) Foreach(f func(uint64) uint64) {
	for _, v := range a {
		f(v)
	}
}

func (a Uint64Array) ForeachIndexed(f func(int, uint64) uint64) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Uint64Array) Reduce(f func(uint64, uint64) uint64) (out uint64) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a Uint64Array) Fold(start uint64, f func(uint64, uint64) uint64) uint64 {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a Uint64Array) ToByteArray() (out ByteArray) {
	for _, v := range a {
		out = append(out, byte(v))
	}
	return
}

func (a Uint64Array) ToFloat32Array() (out Float32Array) {
	for _, v := range a {
		out = append(out, float32(v))
	}
	return
}

func (a Uint64Array) ToFloat64Array() (out Float64Array) {
	for _, v := range a {
		out = append(out, float64(v))
	}
	return
}

func (a Uint64Array) ToIntArray() (out IntArray) {
	for _, v := range a {
		out = append(out, int(v))
	}
	return
}

func (a Uint64Array) ToInt8Array() (out Int8Array) {
	for _, v := range a {
		out = append(out, int8(v))
	}
	return
}

func (a Uint64Array) ToInt16Array() (out Int16Array) {
	for _, v := range a {
		out = append(out, int16(v))
	}
	return
}

func (a Uint64Array) ToInt32Array() (out Int32Array) {
	for _, v := range a {
		out = append(out, int32(v))
	}
	return
}

func (a Uint64Array) ToInt64Array() (out Int64Array) {
	for _, v := range a {
		out = append(out, int64(v))
	}
	return
}

func (a Uint64Array) ToRuneArray() (out RuneArray) {
	for _, v := range a {
		out = append(out, rune(v))
	}
	return
}

func (a Uint64Array) ToStringArray() (out StringArray) {
	for _, v := range a {
		out = append(out, strconv.Itoa(int(v)))
	}
	return
}

func (a Uint64Array) ToUintArray() (out UintArray) {
	for _, v := range a {
		out = append(out, uint(v))
	}
	return
}

func (a Uint64Array) ToUint8Array() (out Uint8Array) {
	for _, v := range a {
		out = append(out, uint8(v))
	}
	return
}

func (a Uint64Array) ToUint16Array() (out Uint16Array) {
	for _, v := range a {
		out = append(out, uint16(v))
	}
	return
}

func (a Uint64Array) ToUint32Array() (out Uint32Array) {
	for _, v := range a {
		out = append(out, uint32(v))
	}
	return
}
