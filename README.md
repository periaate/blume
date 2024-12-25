# blume

## Log
### v0.4.0
- codebase turned into a monorepo.
	- tooling, libraries, and services unified with blume libraries.
- removal of result types
	- patterns `Err[T](args ...any)` and `OK(value)` are still used, but they return `(T, error)`.
	- All `T, error` functions now return `T, error`, and not `Result[T]`.
- option type simplified and streamlined
	- `Option[T]` is now a struct to give more control and less abstraction.
	- `Option[T]` now has only `Must(args ...any)` and `Or()` methods.
- personal formatting style, `blumefmt`, implemented and introduced
	- codebase formatted with `gofmt`.
	- `blumefmt` moved to be purely local by using git filter `clean` and `smudge`.
- structure follows logic:
	- directories at root must be structured such that:
		- the directory can be turned into a git and or go submodule without breaking existing code or workflows.
		- project structure follows locality of behavior.
		- each directory may contain its own `cmd` directory.
	- tools which are not specific to any single library are placed in `cmd`.
- tool changes
	- `tagver`: rewrote, removed bloated dependencies
	- `devious`: renamed `dev` -> `devious`

### v0.3.?+2
- Refactor of option and result semantics, as well as large scale simplifications.
### v0.3.?+1
- Normalization of all libraries to using new higher abstraction patterns.
### v0.3.?
- Refactor from primitive types to type wrappers for strings, arrays, as well as introducing option and result types.


## Licensing
If you can import it, it's MPL 2.0; If you can compile it, it's GPL 3.0.
