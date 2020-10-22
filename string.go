package funct

type StringArray []string

func (a StringArray) Filter(pred func(i string) bool) (out StringArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a StringArray) Foreach(f func(i string) string) {
	for _, v := range a {
		f(v)
	}
}

func (a StringArray) ForeachIndexed(f func(i int, v string) string) {
	for i, v := range a {
		f(i, v)
	}
}

func (a StringArray) Reduce(f func(a, b string) string) (out string) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a StringArray) Fold(start string, f func(a, b string) string) string {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}
