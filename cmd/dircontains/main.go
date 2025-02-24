package main

import . "github.com/periaate/blume"

func main() {
	not := Args().Must()
	Input[String]("pipe").Filter(func(arg String) bool {
		for _, value := range not.Value {
			if value.Contains(arg.String()) {
				return false
			}
		}
		return true
	}).Each(Logs)
}
