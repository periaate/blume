package blume

import "strings"

type Tree[A any] struct {
	This A
	Array[Tree[A]]
}

// Traverse performs a depth-first traversal of the tree and applies the given function to each element
func (t Tree[A]) Traverse(f func(A)) {
	// Apply the function to the current node
	f(t.This)

	// Recursively apply to all children
	for i := range t.Array {
		t.Array[i].Traverse(f)
	}
}

// PrettyPrint creates a string representation of the tree with proper indentation
func (t Tree[A]) PrettyPrint() String {
	return t.prettyPrintWithIndent(-1)
}

// prettyPrintWithIndent is a helper method for pretty printing with the correct indentation level
func (t Tree[A]) prettyPrintWithIndent(indent int) String {
	// Create the indentation string
	var indentStr string
	if indent > 0 {
		indentStr = strings.Repeat("  ", indent)
	}
	var result String

	// Start with the current node
	if P.S(t.This).TrimSpace() != "" {
		result = P.F("%s%v\n", indentStr, t.This)
	}

	// Add all children with increased indentation
	for _, val := range t.Array {
		result += val.prettyPrintWithIndent(indent + 1)
	}

	return result
}

type Composer func(Tree[String]) Tree[String]

var Parser Composer = func(tree Tree[S]) Tree[S] { return tree }

func (parser Composer) Closure(start, end String) Composer {
	return func(tree Tree[String]) Tree[String] {
		tree = parser(tree)
		return tree
	}
}

func (parser Composer) Literal(start, end String) Composer {
	return func(tree Tree[String]) Tree[String] {
		tree = parser(tree)
		return tree
	}
}

func (parser Composer) Define(value, match String) Composer {
	return func(tree Tree[String]) Tree[String] {
		tree = parser(tree)
		return tree
	}
}

// Closure defines embedded closures.
// usage : `Closure("(", ")")`
// input : `ab(10+(3*10))`
// output: `[ab [10+ [3*10]]]`
func Closure(delims ...S) func([]String) Result[Tree[String]] {
	return func(input []S) Result[Tree[S]] {
		delim := make([]Delimiter, 0, len(delims)%2)
		for i := 0; i+1 < len(delims); i += 2 {
			delim = append(delim, Delimiter{delims[i], delims[i+1]})
		}
		return Auto(EmbedDelims(input, delim...))
	}
}

// Literal defines string literals.
// escapes flag is used to define how escaped values are interpreted.
// false means that escapes are ignored.
// "this \"wouldn't\" work" -> ["this ", wouldn't, " work"], but [[this \"would\" work]] -> ["this \"would\" work"]
// true means that escapes are valued, i.e.,   "this \"would\" work"
// usage : `Literal("[[", "]]", false)`
// input : `["abc", "..", "[[Hello", ",", "World", "!]]", "..", "dfg"]`
// output: `["abc", "..", "Hello, World!", "..", "dfg"]`
func Literal(start, end String, escapes bool) func([]String) []String {
	return func(input []S) []S {
		return input
	}
}

// Define builtin functions or operators, inputs are handled accordingly.
// usage : `Define("ADD", "+")`
// input : `["1+2", "+", "3+4"]`
// output: `["1", "ADD", "2", "ADD", "3", "ADD", "4"]`
func Define(value, match String) func([]String) []String {
	return func(input []String) []S {
		return input
	}
}

type Delimiter struct {
	Start String
	End   String
}

func EmbedDelims(sar []String, delims ...Delimiter) Tree[S] {
	car := make([]Tree[S], len(sar))
	for i, s := range sar {
		car[i].This = s
	}
	res, _ := embeds(car, delims)
	return res
}

func embeds(car []Tree[S], delims []Delimiter) (res Tree[S], v int) {
	for i := 0; len(car) > i; i++ {
		v := car[i]
		matched := false
		for _, delim := range delims {
			switch v.This {
			case delim.Start:
				r, k := embeds(car[i+1:], delims)
				i += k
				res.Array = res.Append(r)
				matched = true
			case delim.End:
				return res, i + 1
			}
			if matched {
				break
			}
		}
		if !matched {
			res.Array = res.Append(v)
		}
	}

	return res, 0
}

// func CaptureDelims(str string, keep bool, delims ...rune) (res []string, err error) {
// 	if len(str) == 0 {
// 		err = fmt.Errorf("empty string")
// 		return
// 	}
//
// 	if len(delims) == 0 {
// 		err = fmt.Errorf("no delimiters provided")
// 		return
// 	}
//
// 	if len(delims)%2 != 0 {
// 		err = fmt.Errorf("odd number of delimiters provided")
// 		return
// 	}
//
// 	start := map[rune]rune{}
// 	for i := 0; i < len(delims); i += 2 {
// 		start[delims[i]] = delims[i+1]
// 	}
//
// 	var capturing bool
// 	var end rune
// 	sb := strings.Builder{}
//
// 	for _, r := range str {
//
// 		if capturing {
// 			if r == end {
// 				capturing = false
// 				if keep {
// 					sb.WriteRune(r)
// 				}
// 				res = append(res, sb.String())
// 				sb.Reset()
// 				continue
// 			}
// 			sb.WriteRune(r)
// 			continue
// 		}
//
// 		if v, ok := start[r]; ok {
// 			capturing = true
// 			end = v
// 			res = append(res, sb.String())
// 			sb.Reset()
// 			if keep {
// 				sb.WriteRune(r)
// 			}
// 			continue
// 		}
//
// 		sb.WriteRune(r)
// 	}
//
// 	if sb.Len() > 0 {
// 		res = append(res, sb.String())
// 	}
// 	return
// }
