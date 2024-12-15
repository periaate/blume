package main

import (
	"log"
	"os"

	gf "github.com/jessevdk/go-flags"
)

var Opts Options

var Args []string

func main() {

}

// TODO: Custom pattern files
type Options struct {
	Match bool `long:"match" description:"Perceptually hash the given filenames, printing out all perfect matches."`
}

func init() {
	Opts = Options{}
	args, err := gf.Parse(&Opts)
	if err != nil {
		if gf.WroteHelp(err) {
			os.Exit(0)
		}
		log.Fatalln("Error parsing flags:", err)
	}
	Args = args
}
