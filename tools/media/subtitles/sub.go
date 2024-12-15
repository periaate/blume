package subtitles

import (
	"bytes"
	"fmt"
	"os/exec"
)

type SubtitleResult struct {
	Index         int
	FilePath      string
	SubtitleBytes []byte
}

// ExtractSubtitles takes an array of media file paths and extracts their first subtitle tracks.
func ExtractSubtitles(filePaths []string) ([]SubtitleResult, error) {
	var results []SubtitleResult

	for i, path := range filePaths {
		cmd := exec.Command("ffmpeg", "-i", path, "-map", "0:s:0", "pipe:1.vtt")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("error executing ffmpeg for file %s: %w", path, err)
		}

		results = append(results, SubtitleResult{
			Index:         i,
			FilePath:      path,
			SubtitleBytes: out.Bytes(),
		})
	}

	return results, nil
}
