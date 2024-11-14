# blume
Blume is a cohesive, prioritizing consistency and composability of `blume` libraries.

## primitive libraries
- `blume/gen`: type definitions and the HoFs which define the semantics of `blume`.
- `blume/str`: all necessary tools to work with strings with `blume` semantics.

## standard libraries
- `blume/fsio`: providing normalized filesystem and IO procedures, while remaining fully compatible with `gen` and `str`.
- `blume/clog`: provides a structured, but human readable and formatted logger and printer.

## experimental standard libraries
- `blume/x/hnet`: for reducing the boilerplate of `http` and `net` related loads.
- `blume/x/rfl`: reflection and metaprogramming with `blume` semantics.

## Philosophy
`blume` is based around the concept of "building" or "constructing" procedures with normalized function signatures. To enable this, `blume` semantics take these building blocks before separately from the actual arguments the procedure will act on.

E.g.,
```
data := [2, 3, 7, 8, 9, 11, 12]
predicate := Or(IsEven, Is(7, 11))
filter := Filter(predicate)
filter(data) // [2, 7, 8, 11, 12]
```
Which can be simplified to just
```
Filter(Or(IsEven, Is(7, 11)))(2, 3, 7, 8, 9, 11, 12) // [2, 7, 8, 11, 12]
```
