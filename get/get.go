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
	if err != nil { Err[S](err) }
	res := S(s).Path()
	res.Println(res.IsDir(), res.EnsureSuffix("/"))
	return Ok(res)
}

func Username() Result[S] { return Exec("whoami", CmdOpt.AdoptEnv()).Run().Then(TrimSpace[S]) }
func Hostname() Result[S] { return Exec("hostname", CmdOpt.AdoptEnv()).Run().Then(TrimSpace[S]) }
func Os()       Result[S] { return Exec("uname", CmdOpt.AdoptEnv().Args("-sor")).Run().Then(TrimSpace[S]) }

