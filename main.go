package main

import (
	"github.com/mrap/mosaic/image"
	"log"
)

func main() {
	img, _ := image.ImageFromFile("test/images/sierra-buttes.jpg")
	avg := image.AvgRGBA(img)
	log.Printf("Avg color %+v", avg)
}
