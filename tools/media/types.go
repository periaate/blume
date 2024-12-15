package media

import "path/filepath"

type Kind uint64

const (
	INVALID Kind = 1 << iota
	IMAGE
	VIDEO
	AUDIO
)

var CntMasks = map[Kind][]string{
	IMAGE: {".jpg", ".jpeg", ".png", ".apng", ".gif", ".bmp", ".webp", ".avif", ".jxl", ".tiff"},
	VIDEO: {".mp4", ".m4v", ".webm", ".mkv", ".avi", ".mov", ".mpg", ".mpeg"},
	AUDIO: {".m4a", ".opus", ".ogg", ".mp3", ".flac", ".wav", ".aac"},
}

var ExtToMaskMap = map[string]Kind{}

func Is(kinds ...Kind) func(key string) bool {
	return func(key string) bool {
		key = filepath.Ext(key)
		kind, ok := ExtToMaskMap[key]
		if !ok {
			return ok
		}

		for _, k := range kinds {
			if kind == k {
				return ok
			}
		}

		return false
	}
}

func RegisterMasks(mask Kind, keys ...string) {
	for _, k := range keys {
		ExtToMaskMap[k] |= mask
	}
}

func init() {
	for k, v := range CntMasks {
		RegisterMasks(k, v...)
	}
}

func GetKind(ext string) Kind {
	v, ok := ExtToMaskMap[ext]
	if !ok {
		return INVALID
	}

	return v
}
