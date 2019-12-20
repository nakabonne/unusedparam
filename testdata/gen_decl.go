package testdata

func _(n, m int) int {
	var a = n
	var b = m
	return a + b
}

func _(n, m int) int { // m is unused
	var a = n
	return a
}

func _(n, m int) int { // n,m is unused
	var a int
	return a
}
