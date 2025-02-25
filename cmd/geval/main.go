package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/periaate/blume"
)

func main() {
	var arg String
	iargs := Args().Must()
	switch {
	case iargs.Len().Eq(0):
		panic("not enough arguments")
	case iargs.Len().Gt(1):
		arg = String(fmt.Sprintf("%s(\"%s\")", iargs.Value[0], strings.Join(Map(StoD[String])(iargs.Value[1:]), "\", \"")))
	default:
		arg = iargs.Value[0]
	}
	path := filepath.Join(Must(os.UserHomeDir()), "github.com", "periaate", "blume", "cmd", "geval", "gotemp")
	Must(os.MkdirAll(path, 0755))
	Must(os.Chdir(path))
	filePath := filepath.Join(path, "main.go")
	file := Must(os.Create(filePath))
	defer file.Close()
	Must(fmt.Fprintf(file, "package main\n\nimport . \"github.com/periaate/blume\"\n\nvar _ String\n\nfunc main() {\n\tprintln(%s)\n}", arg))
	Must(file.Sync())
	Must(Exec("go", "mod", "tidy").Silent())
	Must(Exec("goimports", "-w", filePath).Silent())
	Must(Exec("go", "run", filePath).Run())
}
