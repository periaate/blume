package hnet

type ContentType string

const (
	Stream ContentType = "application/octet-stream"
	Text   ContentType = "text/plain"
	JSON   ContentType = "application/json"
	HTML   ContentType = "text/html"
	CSS    ContentType = "text/css"
	JS     ContentType = "text/javascript"
	MP3    ContentType = "audio/mp3"
	OGG    ContentType = "audio/ogg"
	JPEG   ContentType = "image/jpeg"
	PNG    ContentType = "image/png"
	GIF    ContentType = "image/gif"
	MP4    ContentType = "video/mp4"
	WEBM   ContentType = "video/webm"
	MKV    ContentType = "video/mkv"
)
