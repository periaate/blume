package T

type Condition[A any] func(A) Error[string]
