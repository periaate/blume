package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Masterminds/semver/v3"
	"github.com/periaate/blume"
	"github.com/periaate/blume/filter"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/has"
	"github.com/periaate/blume/is"
	"github.com/periaate/blume/pred"
	"github.com/periaate/blume/str"
)

func main() {
	args := blume.Must(fsio.Args(is.NotEmpty[string]))

	switch {
	case pred.Any(is.Any("h", "help"))(args):
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

	args = filter.Any(is.Any("v", "version", "h", "help", "patch", "minor", "major"))(args)

	cmd := exec.Command("git", "tag")
	buf := blume.Buf()
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to fetch tags: %s", err)
	}
	var lastTag *semver.Version

	tags := str.Split(buf.String(), false, "\n")
	for _, t := range tags {
		v, err := semver.NewVersion(t)
		if err == nil && (lastTag == nil || v.GreaterThan(lastTag)) {
			lastTag = v
		}
	}

	if err != nil {
		log.Fatalf("Failed to iterate tags: %s", err)
	}
	if lastTag == nil {
		lastTag, _ = semver.NewVersion("0.0.0")
	}

	switch {
	case has.Any(args...)("major"):
		fmt.Printf("v%s", lastTag.IncMajor())
	case has.Any(args...)("minor"):
		fmt.Printf("v%s", lastTag.IncMinor())
	case has.Any(args...)("patch"):
		fmt.Printf("v%s", lastTag.IncPatch())
	default:
		fmt.Printf("v%s", lastTag)
	}
}
