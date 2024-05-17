package fill

import (
	"image/color"
	"machine"
)

func Rect(x, y, w, h int16, c color.RGBA) {
	// https://www.youtube.com/watch?v=G4aHKNUBAgo
	// Sensacional a explicação de percorrer linha ao invés de coluna.
	// Resolvi deixar aqui compartilhado.
	// Bom para eu não esquecer também.
	// Finalmente eu pude entender porque no projeto do TinyGBA eles fazem assim.
	for i := int16(0); i < h; i++ {
		for j := int16(0); j < w; j++ {
			machine.Display.SetPixel(j+x, i+y, c)
		}
	}
}
