package gen

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/types"
)


func TestSplitWithAll(t *testing.T) {
	tst := `(foo (bar baz abc))`
	delims := []string{"(", ")", " "}
	res := SplitWithAll(tst, true, delims...)
	for i, r := range res {
		fmt.Println(i+1, r)
	}
	if len(res) != 11 { t.Fatalf("expected 6, got %d", len(res)) }
}

func TestEmbed(t *testing.T) {
	cases := []string{
		"(+ 1 2)",
		"(- 10 3 2)",
		"(* 4 5)",
		"(/ 20 4)",
		"(+ (* 2 3) (/ 8 4))",
		"(- (+ 3 5) (* 2 4))",
		"(/ (* (+ 3 2) (- 10 1)) 3)",
	}
	delims := []string{"(", ")", " "}

	for _, testCase := range cases {
		res := SplitWithAll(testCase, true, delims...)

		res = Filter(Not(Is(" ")))(res)

		ebd := EmbedDelims(res, [2]string{"(", ")"})

		fmt.Printf("|||\n")
		traverse(ebd, 0)
	}
	fmt.Printf("|||\n")
}

func traverse(h types.Tree[string], depth int) {
	f := true
	for _, v := range h.Nodes {
		if len(v.Nodes) != 0 {
			f = true
			traverse(v, depth+4)
			continue
		}

		add := " "
		if f {
			add = ">"
			f = false
		}

		fmt.Printf("%s%s\"%s\"\n", strings.Repeat(" ", depth), add, v.Value)
	}
}

func TestEval(t *testing.T) {
	testCase := "(+ 2 (* 2 4) (/ (** 2 8) 8))"
	delims := []string{"(", ")", " "}

	res := SplitWithAll(testCase, true, delims...)

	res = Filter(Not(Is(" ")))(res)

	ebd := EmbedDelims(res, [2]string{"(", ")"})

	traverse(ebd.Nodes[0], 0)

	nres := eval(ebd.Nodes[0])
	fmt.Println(nres)
}

func eval(h types.Tree[string]) string {
	sar := []string{}
	for _, v := range h.Nodes {
		if len(v.Nodes) != 0 {
			found := eval(v)
			sar = append(sar, found)
			fmt.Println("found from eval", found)
			continue
		}
		sar = append(sar, v.Value)
	}

	vals := []int{}
	fmt.Println("0th", h.Nodes[0])

	for _, v := range sar[1:] {
		fmt.Println(v)
		i, _ := strconv.Atoi(v)
		vals = append(vals, i)
	}

	r := sar[0]
	fmt.Println("trying to find op", r, h)
	op, ok := opmap[r]
	if !ok { panic("this isn't supposed to happen") }

	res := vals[0]
	for _, v := range vals[1:] {
		fmt.Println("calling "+r+" with", res, v)
		res = op(res, v)
	}

	fmt.Println("eval", r, "vals", vals, "res", res)
	return fmt.Sprint(res)
}

var opmap = map[string]func(int, int) int{
	"+":  func(a, b int) int { return a + b },
	"-":  func(a, b int) int { return a - b },
	"*":  func(a, b int) int { return a * b },
	"/":  func(a, b int) int { return a / b },
	"%":  func(a, b int) int { return a % b },
	"**": pow, // func(a, b int) int { return int(math.Pow(float64(b), float64(a))) },
}

func pow(a, b int) (res int) {
	res = 1
	for range b {
		res *= a
	}
	return
}
