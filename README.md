# blume

`gen`: Generic, general functions, types, and more.
- Generic, constructed functions, `Filter`, `Map`, etc.
- String functions and utilities.
- Array functions and utilities.
- Map types and wrappers, `Sync`, `Expiring`, etc.
`gen/T`: Interfaces, types, basic implementations, etc. These implementations are trivial, and serve as the "cores" for more complex implementations.
- `Error[any]`: A rich error interface.
- `Result[any]`: A result type, similar to one found in rust.
- `Str[~string]`: A string type, defining numerous methods which make working with strings easier.
- `Tree[any]`: A tree type, providing basic tree functionality.
- `Err[any]`: Implementation of `Error`.
- Various functions for creation of `T` types.
`hnet`: Utilities for working with http.
- `Header`: All valid HTTP headers as constants.
- `Status`: All valid HTTP status codes as constants.
- `URL`: Utilities for working with, validating, manipulating, and using URLs.
- `Request` and `Response`: composable workflows with HTTP requests. (URL -> Request -> Response).
`hnet/auth`: Stateful, userless, session based auth system.
`fsio`: File system related utilities and wrappers.
- Functions for working with IO, e.g., `Args`, `ReadPipe`, `HasPipe`, etc.
- Functions for working with file systems, e.g., `IsDir`, `EnsureDir`, etc.
`yap`: a simple and configurable logging library, focusing on human readability and ease of use.

`tools`: Tools used or developed for `blume`.
`tools/mdn`: Generates the `Header` and `Status` enums for `hent` from the MDN Web Docs.
`tools/gometa`: Macro/code generation program, which supports Go. Used specfically for working around Go's type system limitations by "deriving" "traits" for types, as seen in rust.

`typ`: gometa generated types.
- `String`: gometa generated implementation of `T.Str`.
