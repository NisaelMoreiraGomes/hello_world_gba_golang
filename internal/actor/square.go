package actor

import (
	"image/color"

	"github.com/NisaelMoreiraGomes/hello_world_gba_golang/internal/fill"
)

type Square struct {
	X, Y, Size int16
	C          color.RGBA
}

func (s Square) Draw() {
	fill.Rect(s.X, s.Y, s.Size, s.Size, s.C)
}
