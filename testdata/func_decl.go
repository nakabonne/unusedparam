package testdata

func _(n, m int) int {
	a := func() int {
		return n + m
	}
	return a()
}

func _(n, m int) int { // m is unused
	a := func() int {
		return n
	}
	return a()
}

func _(n, m int) int { // n,m is unused
	a := func() int {
		return 1
	}
	return a()
}
