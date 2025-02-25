package main

import (
	"fmt"
	"strings"

	. "github.com/periaate/blume"
)

/*
git %3?clone git@github.com:%1/%2 ...

search     = pacman -Ss ...
add        = sudo pacman -S ...
update     = sudo pacman -U ...
update all = sudo pacman -Syu ...
remove     = sudo pacman -R ...
=pacman ...
*/

func main() {
	args := Args().Must()
	name := args.Get(0).Must()
	args = ToArray(args.Value[1:])
	file := Read("~", "github.com", "periaate", "blume", "cmd", "query", "testing", name+".cmd").Must()
	var value String
	if file.Contains("=") {
		name := args.Get(0).Or("")
		if args.Len().Gt(1) {
			args = ToArray(args.Value[1:])
		}
		lines := Split(file.String(), false, "\n")
		splits := Map(func(s string) [2]string {
			split := Map(TrimSpace[string])(Split(s, false, "="))
			if len(split) == 1 {
				return [2]string{name.String(), split[0]}
			}
			return [2]string{split[0], split[1]}
		})(lines)
		value = String(ToArray(splits).First(func(value [2]string) bool {
			return name.String() == value[0]
		}).Must()[1])
	} else {
		value = file.Replace("\n", "").TrimSpace()
	}

	i := 0
	for value.Contains(fmt.Sprintf("%%%v", i+1)) {
		i++
		val := args.Get(i - 1)
		if val.IsOk() {
			cur := fmt.Sprintf("%%%v", i)
			reg := fmt.Sprintf("%%%v[?]+([A-z]+ )", i)
			value = value.ReplaceRegex(reg, cur+" ")
			value = value.Replace(cur, val.Value.String())
		} else {
			cur := fmt.Sprintf("%%%v?", i)
			value = value.Replace(cur, "")
		}
	}
	if args.Len().Ge(i) {
		value = value.ReplaceSuffix("...", strings.Join(Map(StoD[String])(args.Value[i:]), " "))
	} else {
		value = value.ReplaceSuffix("...", "")
	}
	Must(Exec("echo", value).Run())
}
