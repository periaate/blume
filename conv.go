package blume

func StoS[A, B ~string](value A) B   { return B(value) }
func ItoI[A, B Numeric](value A) B   { return B(value) }
func StoD[A ~string](value A) string { return string(value) }
