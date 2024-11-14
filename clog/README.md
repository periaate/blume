# clog

```go
import "github.com/periaate/blume/clog"
```

Package clog wraps log/slog with a normalized indent, humanized, and colorized style.

Example:

```
DEBUG @ main.go:111      MSG:<a message>; KEY:<Values here>; err:<nil>;
DEBUG @ main.go:111      MSG:<another message>; KEY:<Values here longer value>; err:<nil>;
DEBUG @ main.go:111      MSG:<a message>;       KEY:<err will be adjusted>;     err:<nil>;
```
