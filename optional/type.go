package optional

type Predicate[T any] func(t T) bool
type Function[T any, V any] func(t T) V
type Consumer[T any] func(t T)
type Supplier[T any] func() T
type Runnable func()
