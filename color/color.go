package color

import (
	"image/color"
)

func HexValue(c color.RGBA) uint32 {
	return uint32(c.A<<24 | c.R<<16 | c.G<<8 | c.B)
}

func NormalizedRGB(c color.RGBA) (r, g, b uint8) {
	return c.R >> 16 & 0xFF,
		c.G >> 8 & 0xFF,
		c.B & 0xFF
}
