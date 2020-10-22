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
