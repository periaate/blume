package core

func LT[N Numeric](n N) Monadic[N, bool]  { return func(val N) bool { return val < n } }
func GT[N Numeric](n N) Monadic[N, bool]  { return func(val N) bool { return val > n } }
func LTE[N Numeric](n N) Monadic[N, bool] { return func(val N) bool { return val <= n } }
func GTE[N Numeric](n N) Monadic[N, bool] { return func(val N) bool { return val >= n } }
func EQ[N Numeric](n N) Monadic[N, bool]  { return func(val N) bool { return val == n } }
func NEQ[N Numeric](n N) Monadic[N, bool] { return func(val N) bool { return val != n } }
func Len[B any, A Lennable[B]](a A) int   { return len(a) }
