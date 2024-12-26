# blume
blume contains Go libraries and software.

## Log
### v0.5.0
- removed `blume/auth`.
- removed `blume/hnet`.
- removed `blume/media`.
- removed `blume/cmd/runn`.
- removed `blume/cmd/filter`.
- removed `blume/cmd/licenser`.
- added new library `blume/pred` for predicates.
	- moved predicate operations from `blume` to `blume/pred`.
	- moved predicates from `blume` to `blume/pred/is` and `blume/pred/has`.
- added new library `blume/types` for types.
	- moved `blume/maps` to `blume/types/maps`.
	- moved string operations from `blume` to `blume/types/str`.
	- moved number related functionality from `blume` to `blume/types/num`.
- `blume/fsio` changes:
	- moved `blume/blob` to `blume/fsio/blob`
	- `blume/fsio/blob` rewritten.
	- introduced `Traverse` function.
		- Breadth first search.
		- Given a walk function which has control over traversal.
		- reimplemented `Find` and `First` to use `Traverse`.
	- removed file path operations.
	- added new library `blume/fsio/ft`:
		- provides unified `Type` type to work with file extensions and mime types
		- provides predicates to check the `Kind` of a `Type`; Video, Audio, Media, Code, ...
- moved `blume/auth/fwauth` to `blume/cmd/fwauth`.
- `blume/types/maps` changes:
	- removed internal type `link`.
	- removed `Expiring` type.
	- introduced `Validated` type.
		- takes a `func(K, V) bool` as an argument.
		- called both on `Get` and `Set` operations.
		- invalid KV pair on `Set` is a noop and returns `false`.
		- invalid KV pair on `Get` removes the pair from the map.
	- introduced `Map` interface.
- `blume` changes:
	- removed everything.
	- introduced `Or[A any](A, A, ...any) A` function.
		- given two values, `A` must `comparable`. The non-zero value is returned. Otherwise,
		- the last value is checked for `bool == true` or `error != nil`.
	- introduced `Must[A any]A, ...any) A` function.
		- the last value is checked for `bool == true` or `error != nil`.
	- introduced `Buf(any) *bytes.Buffer`.
		- takes in anything, and attempts to turn it into a `*bytes.Buffer`.
		- if `string`, turned to bytes to create the buffer.
		- if `[]byte`, used to create the buffer.
		- if `io.Reader`, copies contents to the buffer
		- otherwise return an empty buffer.

### Pre v0.5.0
- Bloat
- Spaghetti
