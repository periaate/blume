package gen

var (
	_                 = Filter(Is(1, 2, 3))
	_                 = All(Is(1, 2, 3))
	_                 = Map(Is(1, 2, 3))
	IsEven            = func(n int) bool { return n%2 == 0 }
	IsPositive        = func(n int) bool { return n > 0 }
	IsEvenAndPositive = And(IsEven, IsPositive)
	_                 = And(IsEven, Or(IsPositive, Is(-20)))
)
