package utils

func IfReturn[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}

	return falseValue
}

func IfThenReturn[T any](condition bool, trueFn func() T, falseFn func() T) T {
	if condition {
		return trueFn()
	}

	return falseFn()
}
