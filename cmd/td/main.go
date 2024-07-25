package main

import (
	"fmt"
	"strings"

	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/core/val"
)

func main() { fmt.Println(Must(val.TimeDate(strings.Join(Args(), " ")))) }
