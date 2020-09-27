package funct

type IntArray []int

func (a IntArray) Filter(f func(int) bool) (out IntArray) {
	for _, v := range a {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}
