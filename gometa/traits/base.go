package traits

import (
	"fmt"
	"runtime"
	"strings"
)

var typeSrc = `"github.com/periaate/blume/T"`

func Implement(packageName string, typeName string, baseType string, traits ...string) string {
	res := fmt.Sprintf("package %s\n\nimport %s\n", packageName, typeSrc)
	for _, trait := range traits {
		switch trait {
		case "String":
			res = fmt.Sprintf("%s\n%s", res, String(typeName, baseType))
		}
	}

	return CanonicalizeNewlines(res)
}

// CanonicalizeNewlines normalizes newlines in a string to match the OS's newline format.
func CanonicalizeNewlines(input string) string {
	var newline string
	switch runtime.GOOS {
	case "windows":
		newline = "\r\n"
	default:
		newline = "\n"
	}

	// Replace both `\r\n` and `\n` with the appropriate newline for the OS
	normalized := strings.ReplaceAll(input, "\r\n", "\n")      // Normalize Windows-style newlines
	normalized = strings.ReplaceAll(normalized, "\n", newline) // Apply OS-specific newlines

	return normalized
}
