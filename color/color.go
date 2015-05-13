package color

import (
	"image/color"
)

func HexValue(c color.RGBA) uint32 {
	return uint32(c.A<<24 | c.R<<16 | c.G<<8 | c.B)
}
