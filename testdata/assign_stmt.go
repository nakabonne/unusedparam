package testdata

func _() int {
	return 0
}

func _(n, m int) int {
	num1 := n
	num2 := m
	return num1 + num2
}

func _(n, m int) int {
	sum := n + m
	return sum
}

func _(n, m int) int { // n,m is unused
	num1 := 1
	num2 := 2
	return num1 + num2
}
