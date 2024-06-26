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
