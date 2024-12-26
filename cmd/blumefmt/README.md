# blumefmt

blumefmt is a formatter for Go. \
blumefmt is not a replacement for gofmt. \
blumefmt is designed to be used with git filters.

blumefmt is made with the assumption that all of the following statements are true.
1. Go is verbose.
2. Verbosity leads to visual clutter.
3. Visual clutter introduces cognitive overhead.
4. Reducing visual clutter reduces cognitive overhead.

## TODO
- [ ] Rewrite.
- [ ] Markdown based test cases.
- [ ] Option flags.
- [ ] Write documentation on setting up and usage of git filters.

## Cases
Test and example cases.

### Inline simple blocks and remove newlines between contiguous simple statements
```go have
package test

func Fn(i int) string {
	if abc == "" {
		return "a"
	}

	if abc == "." {
		return "b"
	}

	for _, r := range abc {
		fmt.Println(r)
	}


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
}


func a(input string) {
	if Pred(input) {
		Do("a")
	}
}

func b(input string) {
	if Pred(input) {
		Do("b")
	}
}

func c() {
	Do("c")
}
```

```go want
package test

func Fn(i int) string {
	if abc == "" { return "a" }
	if abc == "." { return "b" }

	for _, r := range abc {
		fmt.Println(r)
	}

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
}

func a(input string) { if Pred(input) { Do("a") } }
func b(input string) { if Pred(input) { Do("b") } }
func c() { Do("c") }
```

### Inline early returns in simple blocks
```txt options
--early-returns
```

```go have
package main

func main() {
	do.Handle("path", func() {
		res, ok := check()
		if !ok {
			write(ok)
			return
		}
		val, err := with(res)
		if err != nil {
			write(err)
			return
		}
		serve(val)
	})

	do.Handle("path", func() {
		err := valid()
		if err != nil {
			write(err)
			return
		}
		write(true)
	})
}
```

```go want
package main

func main() {
	do.Handle("path", func() {
		res, ok := check()
		if !ok { write(ok); return }
		val, err := with(res)
		if err != nil { write(err); return }
		serve(val)
	})

	do.Handle("path", func() {
		err := valid()
		if err != nil { write(err); return }
		write(true)
	})
}
````

### Align contiguous simple statements
```txt options
--align
```

```go have
package test

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
}

func multi() {
	if do(true) {
		return true
	}
	if abc(a, b, c, d, e) {
		return false
	}
	if a && b && c && d && e {
		return abc(a, b, c, d, e)
	}
}

func a(){
	if do() {
		res()
	}
}

func b(i int) bool {
	if do(i) {
		return res()
	}
}

func c(input string) (bool, error) {
	if does(input) {
		return respond(input)
	}
}
```

```go want
package test

func Fn(i int) string {
	switch {
	case i%2 == 0: return "a"
	case i%3 == 0: return "b"
	case i == 7:   return "c"
	default:       return "d"
	}
}

func multi() {
	if do(true)              { return true }
	if abc(a, b, c, d, e)    { return false }
	if a && b && c && d && e { return abc(a, b, c, d, e) }
}

func a()                           { if do()        { res() } }
func b(i int) bool                 { if do(i)       { return res() } }
func c(input string) (bool, error) { if does(input) { return respond(input) } }
```

## License
blumefmt is licensed under GPL 3.0.
