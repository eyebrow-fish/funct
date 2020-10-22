package funct

type StringArray []string

func (a StringArray) Filter(pred func(string) bool) (out StringArray) {
	for _, v := range a {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

func (a StringArray) Foreach(f func(string) string) {
	for _, v := range a {
		f(v)
	}
}

func (a StringArray) ForeachIndexed(f func(int, string) string) {
	for i, v := range a {
		f(i, v)
	}
}

func (a StringArray) Reduce(f func(string, string) string) (out string) {
	for _, v := range a {
		out = f(out, v)
	}
	return
}

func (a StringArray) Fold(start string, f func(string, string) string) string {
	out := start
	for _, v := range a {
		out = f(out, v)
	}
	return out
}

func (a StringArray) Join() (out string) {
	for _, v := range a {
		out += v
	}
	return
}
