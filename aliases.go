package blume

type Option[A any] = Either[A, bool]
type Result[A any] = Either[A, error]
type Pred[A any] = func(A) bool
type Selector[A any] = func(A) [][]int
