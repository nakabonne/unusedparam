package testdata

func _(n, m int) int {
	return n + m
}

func _(n, m int) int { // n,m is unused
	return 1 + 2
}

func _(n, m int) int { // n,m is unused
	a := 1
	return a
}

func _(n, m int) int { // m is unused
	return n
}
