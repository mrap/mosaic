package color

import (
	"image/color"
)

func HexValue(c color.RGBA) uint32 {
	return uint32(uint32(c.A<<24) | uint32(c.R)<<16 | uint32(c.G)<<8 | uint32(c.B))
}
