package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/periaate/blume/fsio"
)

func main() {
	res, ok := fsio.Args()
	if !ok {
		return
	}

	if len(res) == 0 {
		res = append(res, fsio.ReadPipe()...)
	}

	val := res[0]
	if fsio.Exists(val) {
		res, err := os.ReadFile(val)
		if err != nil {
			return
		}
		val = string(res)
	}

	unixTimestampStr := val
	unixTimestamp, err := strconv.ParseInt(unixTimestampStr, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing timestamp: %v\n", err)
		return
	}

	timestampTime := time.Unix(unixTimestamp, 0)
	duration := time.Since(timestampTime)

	fmt.Println(duration.Truncate(time.Second).String())
}
