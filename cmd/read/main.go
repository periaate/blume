package main

import (
	. "github.com/periaate/blume"
)

func main() {
	Args().Must().Each(func(s String) {
		Read(s).Must().Split("\n").Each(Logs)
	})
}
