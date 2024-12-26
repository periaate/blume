package blumefmt

// TODO: switch to md based test cases.

type testCase struct {
	name    string
	have    string
	desired string
}

var testCases []testCase = []testCase{
	{
		name: "Simple example",
		have: `package test

func Ok(isOk bool) string {
	if Ok {
		return "b"
	}
	return "a"
}`,
		desired: `package test

func Ok(isOk bool) string {
	if Ok { return "b" }
	return "a"
}`,
	},
	{
		name: "Remove unnecessary whitespaces - Example 2",
		have: `package test

func Whitespace(abc string) string {
	if abc == "" {
		return "a"
	}

	if abc == "." {
		return "b"
	}

	for _, r := range abc {
		fmt.Println(r)
	}

	return "c"
}`,
		desired: `package test

func Whitespace(abc string) string {
	if abc == "" { return "a" }
	if abc == "." { return "b" }

	for _, r := range abc {
		fmt.Println(r)
	}

	return "c"
}`,
	},
	{
		name: "Inline switch statement cases if simple - Example 1",
		have: `package test

func Fn(i int) string {
	switch {
	case i%2 == 0:
		return "a"
	case i%3 == 0:
		return "b"
	case i == 7:
		return "c"
	default:
		return "d"
	}
}`,
		desired: `package test

func Fn(i int) string {
	switch {
	case i%2 == 0: return "a"
	case i%3 == 0: return "b"
	case i == 7: return "c"
	default: return "d"
	}
}`,
	},
	{
		name: "Inline switch statement cases if simple - Example 2",
		have: `package test

func Fn(i int) string {
	switch {
	case i%2 == 0:
		fmt.Println(i)
		return "a"
	case i%3 == 0:
		return "b"
	case i == 7:
		return "c"
	default:
		res := "hello"
		res += "world"
		return res
	}
}`,
		desired: `package test

func Fn(i int) string {
	switch {
	case i%2 == 0:
		fmt.Println(i)
		return "a"
	case i%3 == 0: return "b"
	case i == 7: return "c"
	default:
		res := "hello"
		res += "world"
		return res
	}
}`,
	},
	{
		name: "Inline functions if single statement",
		have: `package test

func simple() {
	fmt.Println("...")
}`,
		desired: `package test

func simple() { fmt.Println("...") }`,
	},
	{
		name: "Simple statement whitespace pruning",
		have: `package test

func a() {
	Do("a")
}

func b() {
	Do("b")
}

func c() {
	Do("c")
}

func d() {
	Do("d")
}`,
		desired: `package test

func a() { Do("a") }
func b() { Do("b") }
func c() { Do("c") }
func d() { Do("d") }`,
	},
	{
		name: "Recursive simplification; rules should be applied recursively",
		have: `package test

func a(input string) {
	if Pred(input) {
		Do("a")
	}
}

func b(input string) {
	if Pred(input) {
		Do("b")
	}
}`,
		desired: `package test

func a(input string) { if Pred(input) { Do("a") } }
func b(input string) { if Pred(input) { Do("b") } }`,
	},
}
