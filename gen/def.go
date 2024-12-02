package gen

func ArrayOrDefault[A any](inp []A, def ...A) []A {
	if len(inp) == 0 {
		return def
	}
	return inp
}
