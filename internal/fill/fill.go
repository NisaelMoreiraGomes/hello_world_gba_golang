package fill

import (
	"image/color"
	"machine"
)

func Rect(x, y, w, h int16, c color.RGBA) {
	for i := int16(0); i < w; i++ {
		for j := int16(0); j < h; j++ {
			machine.Display.SetPixel(i+w, j+h, c)
		}
	}
}
