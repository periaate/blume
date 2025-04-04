package blush

import . "github.com/periaate/blume"

var arr = []String{
	`pwd | echo`,
	`echo (pwd)`,
	`echo "$(pwd)"`,
	`0...9 || echo`,
	`0...9 |ab| echo ab`,
	`+ 1 2 3 4`,
	`+ 1 2 3 4 > 2 then echo "hello world" else "not this"`,
}
