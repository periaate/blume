package main

import (
	"fmt"
	"strings"

	. "blume/core"
	"blume/core/val"
)

func main() { fmt.Println(Must(val.TimeDate(strings.Join(Args(), " ")))) }
