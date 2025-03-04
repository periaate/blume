package based

import (
	"slices"
	"strings"

	"github.com/periaate/blume"
)

const Alphabet = `0123456789abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTUVWXY`
const l = len(Alphabet)

func match(c rune) int {
	switch c {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'a':
		return 10
	case 'b':
		return 11
	case 'c':
		return 12
	case 'd':
		return 13
	case 'e':
		return 14
	case 'f':
		return 15
	case 'g':
		return 16
	case 'h':
		return 17
	case 'i':
		return 18
	case 'j':
		return 19
	case 'k':
		return 20
	case 'l':
		return 21
	case 'm':
		return 22
	case 'n':
		return 23
	case 'o':
		return 24
	case 'p':
		return 25
	case 'q':
		return 26
	case 'r':
		return 27
	case 's':
		return 28
	case 't':
		return 29
	case 'u':
		return 30
	case 'v':
		return 31
	case 'w':
		return 32
	case 'x':
		return 33
	case 'y':
		return 34
	case 'A':
		return 35
	case 'B':
		return 36
	case 'C':
		return 37
	case 'D':
		return 38
	case 'E':
		return 39
	case 'F':
		return 40
	case 'G':
		return 41
	case 'H':
		return 42
	case 'I':
		return 43
	case 'J':
		return 44
	case 'K':
		return 45
	case 'L':
		return 46
	case 'M':
		return 47
	case 'N':
		return 48
	case 'O':
		return 49
	case 'P':
		return 50
	case 'Q':
		return 51
	case 'R':
		return 52
	case 'S':
		return 53
	case 'T':
		return 54
	case 'U':
		return 55
	case 'V':
		return 56
	case 'W':
		return 57
	case 'X':
		return 58
	case 'Y':
		return 59
	default:
		return 0
	}
}

func Encode[N blume.Integer](n N) (res string) {
	if n == 0 {
		return string(Alphabet[0])
	}
	num := blume.Abs(n)

	var builder strings.Builder
	for num > 0 {
		remainder := int(num % N(l))
		builder.WriteByte(Alphabet[remainder])
		num = num / N(l)
	}

	bytes := []byte(builder.String())
	slices.Reverse(bytes)
	return string(bytes)
}

func Decode[S ~string](input S) (sum int) {
	for i, char := range input {
		if i > 0 {
			sum *= l
		}
		sum += match(char)
	}
	return
}
