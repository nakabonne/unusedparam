package testdata

func _(n, m int) int {
	go func() int {
		return n + m
	}()
	return 0
}

func _(n, m int) int { // m is unused
	go func(i int) int {
		return i
	}(n)
	return 0
}

func _(n, m int) int { // n,m is unused
	go func(i int) int {
		return i
	}(1)
	return 0
}
