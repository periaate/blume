package traits

import (
	"fmt"
	"runtime"
	"strings"
)

var typeSrc = `import (
	"strings"

	"github.com/periaate/blume/gen"
)`

func Implement(packageName string, name string, base string, traits ...string) string {
	fmt.Println(packageName)
	fmt.Println(name)
	fmt.Println(base)
	fmt.Println(traits[0])
	res := fmt.Sprintf("package %s\n\n%s\n", packageName, typeSrc)
	for _, trait := range traits {
		switch trait {
		case "String": res = fmt.Sprintf("%s\n%s", res, String(name, base))
		}
	}

	return CanonicalizeNewlines(res)
}

// CanonicalizeNewlines normalizes newlines in a string to match the OS's newline format.
func CanonicalizeNewlines(input string) string {
	var newline string
	switch runtime.GOOS {
	case "windows": newline = "\r\n"
	default: newline = "\n"
	}

	// Replace both `\r\n` and `\n` with the appropriate newline for the OS
	normalized := strings.ReplaceAll(input, "\r\n", "\n")      // Normalize Windows-style newlines
	normalized = strings.ReplaceAll(normalized, "\n", newline) // Apply OS-specific newlines

	return normalized
}
