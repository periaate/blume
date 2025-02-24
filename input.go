package blume

import "os"

func Input[S ~string](from ...string) Array[String] {
	res := []String{}
	for _, arg := range from {
		switch String(arg).ToLower() {
		case "args":
			res = append(res, Args().Must().Value...)
		case "pipe", "piped":
			res = append(res, Piped(os.Stdin).Must().Value...)
		}
	}

	return ToArray(res)
}
