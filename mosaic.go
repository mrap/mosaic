package mosaic

import (
	"image"
	"log"

	. "github.com/mrap/mosaic/image"
)

type Mosaic struct {
	BaseImage image.Image
	Images    []image.Image
}

func NewMosaic(baseImgPath string, imgPaths []string) (*Mosaic, error) {
	m := new(Mosaic)

	var err error
	m.BaseImage, err = ImageFromFile(baseImgPath)
	if err != nil {
		log.Println("NewMosaic error when loading base image:", err)
		return m, err
	}

	var img image.Image
	for _, p := range imgPaths {
		img, err = ImageFromFile(baseImgPath)
		if err != nil {
			log.Println("NewMosaic error when loading image path:", p, err)
		} else {
			m.Images = append(m.Images, img)
		}
	}

	return m, nil
}
