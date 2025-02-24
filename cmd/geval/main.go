package main

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/periaate/blume"
)

func main() {
	path := filepath.Join(Must(os.UserHomeDir()), "go", "src", "gotemp")
	Must(os.MkdirAll(path, 0755))
	Must(os.Chdir(path))
	filePath := filepath.Join(path, "main.go")
	file := Must(os.Create(filePath))
	defer file.Close()
	Must(fmt.Fprintf(file, "package main\n\nimport . \"github.com/periaate/blume\"\n\nvar _ String\n\nfunc main() {\n\tprintln(%s)\n}", os.Args[1]))
	Must(file.Sync())
	Must(Exec("go", "mod", "tidy").Silent())
	Must(Exec("goimports", "-w", filePath).Silent())
	Must(Exec("go", "run", filePath).Run())
}
