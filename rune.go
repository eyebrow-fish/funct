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
