package str

func IsDigit(str string) bool {
	for _, r := range str {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			return false
		}
	}
	return true
}

func IsNumber(str string) bool {
	if HasPrefix("-", "+")(str) {
		str = Shift(1)(str)
	}

	for _, r := range str {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ',':
			continue
		default:
			return false
		}
	}
	return true
}
