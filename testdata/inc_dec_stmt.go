package testdata

func _(n, m int) int {
	n++
	m--
	return 0
}

func _(n, m int) int { // m is unused
	n++
	return 0
}

func _(n, m int) int { // n,m is unused
	var a int
	a++
	return a
}
