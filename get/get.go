package get

import (
	"os"

	. "github.com/periaate/blume"
)

func IP() String {
	s := Exec("fish", CmdOpt.Args("-Nc", "ip -4 addr | grep eno1")).Run().OrExit().TrimSpace().Replace("\n", " ")
	res := ReplaceRegex(".*inet ([0-9|\\.]*)/.*", "$1")(s)
	return res
}

func PublicIP() Result[S] { return S("https://api.ipify.org").Request() }

func Wd() Result[S] {
	s, err := os.Getwd()
	return Auto(S(s), err)
}

func Username() Result[S] { return Exec("whoami", CmdOpt.AdoptEnv()).Run() }
func Hostname() Result[S] { return Exec("hostname", CmdOpt.AdoptEnv()).Run() }

