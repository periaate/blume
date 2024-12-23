package media

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"golang.org/x/image/draw"
)

func ImageFromReader(r io.Reader) (img image.Image, err error) {
	img, _, err = image.Decode(r)
	return img, err
}

/*
CapImageSize creates an empty RGBA image with the same aspect ration as the input,
but capped to the maxSize provided,
such that neither height nor width are above that value.

edge cases:
  - source file very tall, very wide, resulting in an unusable image or a very large filesize
*/
func CapImageSize(img image.Image, maxSize int) *image.RGBA {
	srcWidth := img.Bounds().Dx()
	srcHeight := img.Bounds().Dy()
	newWidth, newHeight := srcWidth, srcHeight

	if srcWidth > maxSize || srcHeight > maxSize {
		if srcWidth > srcHeight {
			newWidth = maxSize
			newHeight = (srcHeight * maxSize) / srcWidth
		} else {
			newHeight = maxSize
			newWidth = (srcWidth * maxSize) / srcHeight
		}
	}

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	return dst
}

// MinImageSize cap the smaller dimension to maxSize while maintaining aspect ration.
func MinImageSize(img image.Image, maxSize int) *image.RGBA {
	srcWidth := img.Bounds().Dx()
	srcHeight := img.Bounds().Dy()
	newWidth, newHeight := srcWidth, srcHeight

	if srcWidth > maxSize || srcHeight > maxSize {
		if srcWidth < srcHeight {
			newWidth = maxSize
			newHeight = (srcHeight * maxSize) / srcWidth
		} else {
			newHeight = maxSize
			newWidth = (srcWidth * maxSize) / srcHeight
		}
	}

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	return dst
}

// Draw copies src to dst, using either a fast algorithm, or a quality algorithm.
func Draw(src image.Image, dst *image.RGBA, quality bool) {
	switch quality {
	case true:
		draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)
	case false:
		draw.ApproxBiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)
	}
}

func ReduceImageSize(img image.Image, maxDimension int) *image.RGBA {
	srcWidth := img.Bounds().Dx()
	srcHeight := img.Bounds().Dy()
	newWidth, newHeight := srcWidth, srcHeight

	if srcWidth > maxDimension || srcHeight > maxDimension {
		if srcWidth > srcHeight {
			newWidth = maxDimension
			newHeight = (srcHeight * maxDimension) / srcWidth
		} else {
			newHeight = maxDimension
			newWidth = (srcWidth * maxDimension) / srcHeight
		}
	}

	// Create a new image with the calculated dimensions
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	// Scale the image using the ApproxBiLinear interpolation algorithm
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Src, nil)

	return dst
}

func FlushImage(tar io.Writer, img image.Image, quality int) error {
	return jpeg.Encode(tar, img, &jpeg.Options{Quality: quality})
}
