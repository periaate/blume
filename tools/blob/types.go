package blob

import (
	"strings"

	"github.com/periaate/blume/gen"
)

// Web-safe Base64 alphabet
const webSafeBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

// IntToWebBase64 converts an integer to a web-safe Base64 string with up to two characters.
func Encode(n int) string {
	n = gen.Clamp(0, 4095)(n)

	var result strings.Builder

	// Calculate the two "digits" of the base-64 representation.
	firstDigit := n / 64
	secondDigit := n % 64

	// Append the characters corresponding to the digits.
	if firstDigit > 0 {
		result.WriteByte(webSafeBase64[firstDigit])
	}
	result.WriteByte(webSafeBase64[secondDigit])

	return result.String()
}

type ContentType int

func (c ContentType) Value() int { return int(c) }
func (c ContentType) Fmt() string {
	if c < 64 {
		return "A" + Encode(int(c))
	}
	return Encode(int(c))
}

func (c ContentType) String() string {
	switch c {
	case STREAM:
		return "application/octet-stream"
	case PLAIN:
		return "text/plain"
	case HTML:
		return "text/html"
	case JSON:
		return "application/json"
	case CSS:
		return "text/css"
	case JAVASCRIPT:
		return "text/javascript"
	case MP3:
		return "audio/mp3"
	case OGG:
		return "audio/ogg"
	case JPEG:
		return "image/jpeg"
	case PNG:
		return "image/png"
	case GIF:
		return "image/gif"
	case MP4:
		return "video/mp4"
	case WEBM:
		return "video/webm"
	case MKV:
		return "video/mkv"
	default:
		return "application/octet-stream"
	}
}

const (
	STREAM ContentType = iota
	PLAIN
	HTML
	JSON
	CSS
	JAVASCRIPT
	MP3
	OGG
	JPEG
	PNG
	GIF
	MP4
	WEBM
	MKV
)

func GetCT(str string) ContentType {
	switch str {
	case "application/octet-stream", "AA":
		return STREAM
	case "text/plain", "AB":
		return PLAIN
	case "text/html", "AC":
		return HTML
	case "application/json", "AD":
		return JSON
	case "text/css", "AE":
		return CSS
	case "text/javascript", "AF":
		return JAVASCRIPT
	case "audio/mp3", "AG":
		return MP3
	case "audio/ogg", "AH":
		return OGG
	case "image/jpeg", "AI":
		return JPEG
	case "image/png", "AJ":
		return PNG
	case "image/gif", "AK":
		return GIF
	case "video/mp4", "AL":
		return MP4
	case "video/webm", "AM":
		return WEBM
	case "video/mkv", "AN":
		return MKV
	default:
		return -1
	}
}
