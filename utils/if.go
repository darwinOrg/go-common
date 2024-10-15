package utils

func IfReturn[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}

	return falseValue
}

func IfConsume(condition bool, trueFn func(), falseFn func()) {
	if condition {
		trueFn()
	}

	falseFn()
}

func IfConsumeReturn[T any](condition bool, trueFn func() T, falseFn func() T) T {
	if condition {
		return trueFn()
	}

	return falseFn()
}
