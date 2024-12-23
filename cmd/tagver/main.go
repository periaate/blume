package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Masterminds/semver/v3"
	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/yap"
)

func up(fp string) string { return fsio.Clean(filepath.Dir(fp)) }

func main() {
	args := fsio.Args[string]().Must().Val
	fp := "."

	switch {
	case Any(Is("h", "help"))(args):
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

	sargs := Filter(Not(Is("v", "version", "h", "help", "patch", "minor", "major")))(args)
	if len(sargs) > 0 { fp = sargs[0] }

	args = Filter(Is("v", "version", "h", "help", "patch", "minor", "major"))(args)

	fp = Must(filepath.Abs(fp))
	for fp != "/" {
		if !fsio.Exists(filepath.Join(fp, ".git")) {
			fp = up(fp)
			continue
		}
		break
	}

	tags, err := ex("git", "tag")
	if err != nil { log.Fatalf("Failed to fetch tags: %s", err) }

	var lastTag *semver.Version

	for _, t := range tags.Val {
		v, err := semver.NewVersion(t)
		if err != nil {
			yap.Error("couldn't pares semantic version", "ver", t, "err", err)
			continue
		}
		if lastTag == nil || v.GreaterThan(lastTag) { lastTag = v }
	}

	if err != nil { log.Fatalf("Failed to iterate tags: %s", err) }

	if lastTag == nil { lastTag, _ = semver.NewVersion("0.0.0") }

	switch {
	case Any(Is("major"))(args): fmt.Printf("v%s", lastTag.IncMajor())
	case Any(Is("minor"))(args): fmt.Printf("v%s", lastTag.IncMinor())
	case Any(Is("patch"))(args): fmt.Printf("v%s", lastTag.IncPatch())
	default: fmt.Printf("v%s", lastTag)
	}
}

func ex(comd string, args ...string) (Array[string], error) {
	cmd := exec.Command(comd, args...)
	buf := Buf()
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil { return Err[Array[string]]("error running command err: {:s}", err) }

	return Ok(ToArray(Split(buf.String(), false, "\n")))
}
