package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/gen"
)

func up(fp string) string {
	return fsio.Clean(filepath.Dir(fp))
}

func main() {
	args := fsio.Args()
	fp := "."

	switch {
	case gen.Any(gen.Is("h", "help"))(args...):
		fmt.Println("tagver")
		fmt.Println("tagver is a simple tool to manage semantic versioning tags in git repositories.")
		fmt.Println("Usage:")
		fmt.Println("  tagver [options] [path]")
		fmt.Println("Options:")
		fmt.Println("  h, help    Show this help message.")
		fmt.Println("  v, version Show the current version. Default")
		fmt.Println("  patch      Increment the patch version.")
		fmt.Println("  minor      Increment the minor version.")
		fmt.Println("  major      Increment the major version.")
		os.Exit(0)
	}

	sargs := gen.Filter(gen.Isnt("v", "version", "h", "help", "patch", "minor", "major"))(args)
	if len(sargs) > 0 {
		fp = sargs[0]
	}

	args = gen.Filter(gen.Is("v", "version", "h", "help", "patch", "minor", "major"))(args)

	fp = gen.Must(filepath.Abs(fp))
	for fp != "/" {
		if !fsio.Exists(filepath.Join(fp, ".git")) {
			fp = up(fp)
			continue
		}
		break
	}

	r, err := git.PlainOpen(fp)
	if err != nil {
		log.Fatalf("Failed to open repository: %s", err)
	}

	tagRefs, err := r.Tags()
	if err != nil {
		log.Fatalf("Failed to fetch tags: %s", err)
	}

	var lastTag *semver.Version

	err = tagRefs.ForEach(func(t *plumbing.Reference) error {
		v, err := semver.NewVersion(t.Name().Short())
		if err == nil {
			if lastTag == nil || v.GreaterThan(lastTag) {
				lastTag = v
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to iterate tags: %s", err)
	}

	if lastTag == nil {
		lastTag, _ = semver.NewVersion("0.0.0")
	}

	switch {
	case gen.Any(gen.Is("major"))(args...):
		fmt.Printf("v%s", lastTag.IncMajor())
	case gen.Any(gen.Is("minor"))(args...):
		fmt.Printf("v%s", lastTag.IncMinor())
	case gen.Any(gen.Is("patch"))(args...):
		fmt.Printf("v%s", lastTag.IncPatch())
	default:
		fmt.Printf("v%s", lastTag)
	}
}
