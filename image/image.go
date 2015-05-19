package image

import (
	"image"
	"image/jpeg"
	"os"
	"path"
)

func ImageFromFile(filePath string) (image.Image, error) {
	img, _ := os.Open(filePath)
	defer img.Close()

	switch path.Ext(filePath) {
	case ".jpeg":
		return jpeg.Decode(img)
	default:
		img, _, err := image.Decode(img)
		return img, err
	}
}

func PixelsCount(img image.Image) uint {
	return uint((img.Bounds().Max.X - img.Bounds().Min.X) * (img.Bounds().Max.Y - img.Bounds().Min.Y))
}
