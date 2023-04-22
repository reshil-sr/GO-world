package sum

func Sums(vs ...int) int {
	return sums(vs)
}

func sums(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return sums(vs[1:]) + vs[0]
}
