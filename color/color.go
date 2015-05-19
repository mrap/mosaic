package color

import (
	"image/color"
)

const MaxRGBAValue uint64 = 0xFFFFFFFF
const MaxRGBValue uint32 = 0xFFFFFF

func HexValue(c color.RGBA) uint32 {
	return uint32(c.A)<<24 | uint32(c.R)<<16 | uint32(c.G)<<8 | uint32(c.B)
}

func RGBHexValue(c color.RGBA) uint32 {
	return HexValue(c) & MaxRGBValue
}
