package blume

// import (
// 	"fmt"
// 	"testing"
// )
//
// func TestIsFormat(t *testing.T) {
// 	testCases := []struct {
// 		input  string
// 		expect bool
// 	}{
// 		{"Hello, %s!", true},
// 		{"The value is %d.", true},
// 		{"A float: %f", true},
// 		{"Scientific: %e", true},
// 		{"Boolean: %t", true},
// 		{"A simple string with no formats.", false},
// 		{"This string has an escaped percent sign: 100%%.", false},
// 		{"Value: %v", true},
// 		{"Go-syntax: %#v", true},
// 		{"Type: %T", true},
// 		{"Pointer: %p", true},
// 		{"Another string with multiple formats: %s and %d", true},
// 		{"An escaped percent %% followed by a format %s", true},
// 		{"", false},
// 		{"%", true},
// 		{"%%", false},
// 		{"a%b", true},
// 	}
//
// 	for _, test := range testCases {
// 		t.Run(test.input, func(t *testing.T) { if IsFormat(test.input) != test.expect { t.Fail() } })
// 	}
// }
