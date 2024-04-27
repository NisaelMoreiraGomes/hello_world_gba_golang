package main

import (
	"machine"

	"github.com/NisaelMoreiraGomes/hello_world_gba_golang/internal/fill"
)

func main() {
	machine.Display.Configure()

	var screenWidth, screenHeight = machine.Display.Size()

	var rectW, rectH int16 = 10, 10
	var rectX, rectY = screenWidth/2 - rectW, screenHeight/2 - rectH

	fill.Rect(rectX, rectY, rectW, rectH, fill.Blue)
}
