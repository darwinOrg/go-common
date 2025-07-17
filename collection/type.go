package dgcoll

type Predicate[T any] func(t T) bool
type Function[T any, V any] func(t T) V
type BiFunction[T1 any, T2 any, V any] func(t1 T1, t2 T2) V
type Less[T any] func(t1 T, t2 T) bool

func Identity[T any](t T) T { return t }
