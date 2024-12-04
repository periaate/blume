package hnet

type ContentType string

func (c ContentType) String() string   { return string(c) }
func (c ContentType) Tuple() [2]string { return Content_Type.Tuple(c.String()) }

const (
	Application_Stream ContentType = "application/octet-stream"
	Application_JSON   ContentType = "application/json"
	Text_Plain         ContentType = "text/plain"
	Text_HTML          ContentType = "text/html"
	Text_CSS           ContentType = "text/css"
	Text_JS            ContentType = "text/javascript"
	Audio_MP3          ContentType = "audio/mp3"
	Audio_OGG          ContentType = "audio/ogg"
	Image_JPEG         ContentType = "image/jpeg"
	Image_PNG          ContentType = "image/png"
	Image_GIF          ContentType = "image/gif"
	Video_MP4          ContentType = "video/mp4"
	Video_WEBM         ContentType = "video/webm"
	Video_MKV          ContentType = "video/mkv"
)
