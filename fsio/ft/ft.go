package ft

import "github.com/periaate/blume/pred/has"

type Kind int

const (
	Other Kind = iota
	Image
	Video
	Audio
	Code
	Config
)

type Type struct {
	name   string
	kind   Kind
	header string
	ext    string
}

func (t Type) ContentHeader() string { return t.header }
func (t Type) Ext() string           { return t.ext }

var (
	jpeg       = Type{"jpeg", Image, "image/jpeg", ".jpeg"}
	png        = Type{"png", Image, "image/png", ".png"}
	gif        = Type{"gif", Image, "image/gif", ".gif"}
	apng       = Type{"apng", Image, "image/apng", ".apng"}
	avif       = Type{"avif", Image, "image/avif", ".avif"}
	webp       = Type{"webp", Image, "image/webp", ".webp"}
	jxl        = Type{"jxl", Image, "image/jxl", ".jxl"}
	mp4        = Type{"mp4", Video, "video/mp4", ".mp4"}
	mpeg       = Type{"mpeg", Video, "video/mpeg", ".mpeg"}
	webm       = Type{"webm", Video, "video/webm", ".webm"}
	mkv        = Type{"mkv", Video, "video/x-matroska", ".mkv"}
	mov        = Type{"mov", Video, "video/quicktime", ".mov"}
	ogg        = Type{"ogg", Audio, "audio/ogg", ".ogg"}
	mp3        = Type{"mp3", Audio, "audio/mpeg", ".mp3"}
	opus       = Type{"opus", Audio, "audio/opus", ".opus"}
	text       = Type{"text", Other, "text/plain", ".txt"}
	javascript = Type{"javascript", Code, "application/javascript", ".js"}
	typescript = Type{"typescript", Code, "application/typescript", ".ts"}
	css        = Type{"css", Code, "text/css", ".css"}
	html       = Type{"html", Code, "text/html", ".html"}
	xml        = Type{"xml", Code, "application/xml", ".xml"}
	json       = Type{"json", Config, "application/json", ".json"}
	raw        = Type{"raw", Other, "application/octet-stream", ""}
)

var ctToType map[string]Type = map[string]Type{
	"image/jpeg":               jpeg,
	"image/png":                png,
	"image/gif":                gif,
	"image/apng":               apng,
	"image/avif":               avif,
	"image/webp":               webp,
	"image/jxl":                jxl,
	"video/mp4":                mp4,
	"video/mpeg":               mpeg,
	"video/webm":               webm,
	"video/x-matroska":         mkv,
	"video/quicktime":          mov,
	"audio/ogg":                ogg,
	"audio/mpeg":               mp3,
	"audio/opus":               opus,
	"text/plain":               text,
	"application/javascript":   javascript,
	"application/typescript":   typescript,
	"text/css":                 css,
	"text/html":                html,
	"application/xml":          xml,
	"application/json":         json,
	"application/octet-stream": raw,
}

var extToType map[string]Type = map[string]Type{
	".jpeg": jpeg,
	".jpg":  jpeg, // For common alias
	".png":  png,
	".gif":  gif,
	".apng": apng,
	".avif": avif,
	".webp": webp,
	".jxl":  jxl,
	".mp4":  mp4,
	".mpeg": mpeg,
	".webm": webm,
	".mkv":  mkv,
	".mov":  mov,
	".ogg":  ogg,
	".mp3":  mp3,
	".opus": opus,
	".txt":  text,
	".js":   javascript,
	".ts":   typescript,
	".css":  css,
	".html": html,
	".xml":  xml,
	".json": json,
	"":      raw, // Catch-all for raw
}

func Jpeg() Type       { return jpeg }
func Png() Type        { return png }
func Gif() Type        { return gif }
func Apng() Type       { return apng }
func Avif() Type       { return avif }
func Webp() Type       { return webp }
func Jxl() Type        { return jxl }
func Mp4() Type        { return mp4 }
func Mpeg() Type       { return mpeg }
func Webm() Type       { return webm }
func Mkv() Type        { return mkv }
func Mov() Type        { return mov }
func Ogg() Type        { return ogg }
func Mp3() Type        { return mp3 }
func Opus() Type       { return opus }
func Text() Type       { return text }
func JavaScript() Type { return javascript }
func TypeScript() Type { return typescript }
func Css() Type        { return css }
func Html() Type       { return html }
func Xml() Type        { return xml }
func Json() Type       { return json }
func Raw() Type        { return raw }
func FromContentHeader(str string) (res Type, ok bool) {
	res, ok = ctToType[str]
	return
}

func FromExt(str string) (res Type, ok bool) {
	if len(str) == 0 {
		return
	}
	if !has.Prefix(".")(str) {
		str = "." + str
	}
	res, ok = ctToType[str]
	return
}
