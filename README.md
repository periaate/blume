# blume
An experimental, pragmatic approach to declarative programming in Go.

## Usage
Installation
```
go get github.com/periaate/blume
```

- blume is based on the idea of "construction". A form of function composition which effectively reverses the order in which you would traditionally give arguments.
```go
func main() {
	// we can create a function which checks whether a given `comparable` type is in the arguments.
	myPredicate := blume.Is("Hello", ", ", "World", "!")
	fmt.Println(myPredicate("hi")) // false
	fmt.Println(myPredicate("hello")) // false
	fmt.Println(myPredicate("Hello")) // true
	// this enables us to construct functionality declarative and then utilize it procedurally.
	// for example, if we wanted a filter which must have an https:// prefix
	// and contain `github.com` and `blume`, we can construct it like so:
	myFilter := Filter(
		blume.HasPrefix("https://"),
		// we can use the `blume.PredAnd` to join predicates to a single one with a logical AND
		blume.PredAnd( 
			blume.Contains("github.com"),
			blume.Contains("blume"),
		),
	)

	// blume.Filter is an array function, and acts on array inputs.
	// for single value inputs, use a predicate instead.
	// Most functions in blume do not work in-place. In-place variants might be added in later versions.
	res := myFilter([]string{
		"github.com/periaate/blume",
		"http://github.com/periaate/blume",
		"https://github.com/periaate/blob",
		"https://github.com/periaate/blume",
	})
	fmt.Println(res, len(res)) // ["https://github.com/periaate/blume"] 1
}
```

- blume has many array and string specific functions. It also provides fluent interfaces for these. All functionality base `blume` provides is usable with primitive types without need to use the custom types blume provides.
```go
func main() {
	// blume.String is a string alias with various added methods.
	var str blume.String
	// blume.Array is a struct that wraps a slice, providing blume functions as methods.
	var arr blume.Array[blume.String]
	stringSlice := []string{
		"Hello",
		",",
		" ",
		"World",
		"!",
	}

	// we can create a `blume.Array` from any type by calling `ToArray` with a slice
	myArray := ToArray(stringSlice)

	// blume also provides some utilites to work with type aliases
	// such as `StoS`, which transforms between string aliases
	arr = Map[string, blume.String](StoS)(myArray)
	// blume.Array is not an interface type. We can access the slice inside. This may change.
	fmt.Println(strings.Join(arr.Val, ""))
	// alternatively, we can also call the `Values()` method
	fmt.Println(strings.Join(arr.Values(), ""))
}
```

- blume changes many fundamental patterns of Go, such as error handling.
```go
func main() {
	// blume functions always return a single value or none.
	// return type (T, error) becomes Result[T], while (T, ok) becomes Option[T].
	res := fsio.ReadDir("./") // blume.Result[blume.Array[blume.String]]
	// as this error is unrecoverable there's no need to manually handle it.
	arr := res.Unwrap() // blume.Array[blume.String]
	// package fsio also provides various file system specific predicates, such as `fsio.IsDir`.
	filtered := result.Filter(fsio.IsDir)
}
```

For a more complex example, I will use my build automatization tool to demonstrate making a multifaceted problem into a simple one.
```go
// note: this example is using `github.com/periaate/blume` as a dot import, hence `blume.` isn't used.
func main() {
	args := fsio.Args[string, String](func(s []string) bool { return len(s) >= 1 }).Unwrap()
	arg := args.Shift().Unwrap()

	// fsio.FindFirst recursively (BFS) looks through the given directory
	// returning the first match, if found.
	found := fsio.FindFirst("my/projects/dir/",
		Not(Contains("/.", "node_modules", "target", "build", "data", "Modules", "mpv.net")),
		fsio.IsDir,
		func(f String) bool { return fsio.Base(f) == arg },
	).Unwrap()
	entryOpt := fsio.ReadDir(found).Unwrap().First(IsEntry) // try to find entry file, e.g., main.go
	var entry String
	if entryOpt.Ok() { entry = entryOpt.Unwrap() } // not all languages need entry files

	// we will ascend, trying to find a project root, e.g., contains `go.mod`, `cargo.toml`, etc.
	root := fsio.Ascend(found, IsProject[String]).Unwrap()
	// get the directory and normalize the directory
	root = fsio.Clean(fsio.Dir(root))
	...
}


```
