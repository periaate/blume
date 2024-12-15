package media

import (
	"image"
	"image/color"
)

// GeneratePhash returns a basic perceptual hash of the given image.
func GeneratePhash(img image.Image, scale int) []byte {
	resizedImg := ResizeImage(img, scale, scale)
	Draw(img, resizedImg, false)

	grayImg := toGrayscale(resizedImg)
	avgLight := averageBrightness(grayImg)

	// width * height
	hash := make([]byte, scale*scale/8)
	for i, pixel := range grayImg.Pix {
		if pixel >= avgLight {
			hash[i/8] |= 1 << uint(i%8)
		} else {
			hash[i/8] &= ^(1 << uint(i%8))
		}
	}

	return hash
}

// ResizeImage resizes the given image to the specified width and height.
func ResizeImage(img image.Image, width, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// toGrayscale converts the given image to grayscale.
func toGrayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayImg.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}

	return grayImg
}

// averageBrightness returns the average brightness of the given grayscale image.
func averageBrightness(img *image.Gray) uint8 {
	var total float64
	bounds := img.Bounds()
	area := bounds.Dx() * bounds.Dy()
	for pixel := range img.Pix {
		total += float64(pixel) / float64(area)
	}

	return uint8(total)
}
