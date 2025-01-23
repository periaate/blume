package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/periaate/blume/fsio"
)

func main() {
	res, ok := fsio.Args()
	if !ok || len(res) == 0 {
		res = append(res, "0h0m0s")
	}

	durationStr := res[0]

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Printf("Error parsing duration: %v\n", err)
		return
	}

	// Calculate the target Unix timestamp
	currentTime := time.Now()
	targetTime := currentTime.Add(-duration)
	unixTimestamp := targetTime.Unix()

	fmt.Print(strconv.FormatInt(unixTimestamp, 10))
}
