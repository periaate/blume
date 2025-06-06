package blume

func Gt[N Numeric](arg N) Pred[N] { return func(n N) bool { return n > arg } }
func Ge[N Numeric](arg N) Pred[N] { return func(n N) bool { return n >= arg } }
func Lt[N Numeric](arg N) Pred[N] { return func(n N) bool { return n < arg } }
func Le[N Numeric](arg N) Pred[N] { return func(n N) bool { return n <= arg } }
func Eq[K comparable](arg K) Pred[K] { return func(n K) bool { return n == arg } }
func Ne[K comparable](arg K) Pred[K] { return func(n K) bool { return n != arg } }
