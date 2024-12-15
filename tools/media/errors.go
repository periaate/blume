package media

import "errors"

var (
	ErrNoMIME          = errors.New("no mime could be detected")
	ErrUnsupportedMIME = errors.New("detected mime is unsupported")
	ErrFFMPEGNotFound  = errors.New("ffmpeg not found in PATH")
)
