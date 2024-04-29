package main

import (
	"machine"

	"github.com/NisaelMoreiraGomes/hello_world_gba_golang/internal/actor"
	"github.com/NisaelMoreiraGomes/hello_world_gba_golang/internal/fill"
	"github.com/NisaelMoreiraGomes/hello_world_gba_golang/internal/pads"
	"github.com/NisaelMoreiraGomes/hello_world_gba_golang/internal/timing"
)

var player = actor.Square{0, 0, 20, fill.Blue}
var shadow = actor.Square{0, 0, 20, fill.Black}
var moved = true

func main() {
	machine.Display.Configure()

	screenWidth, screenHeight := machine.Display.Size()
	centerX, centerY := screenWidth/2-20, screenHeight/2-20

	player.X = centerX
	player.Y = centerY

	shadow.X = centerX
	shadow.Y = centerY

	fill.Rect(0, 0, screenWidth, screenHeight, fill.Black)

	for {
		timing.WaitForVBlank()
		update()
		draw()
	}
}

func update() {
	key := pads.ReadPad()

	switch {
	case pads.Up.IsPressed(key):
		shadow.Y = player.Y
		player.Y -= 1
		moved = true
	case pads.Down.IsPressed(key):
		shadow.Y = player.Y
		player.Y += 1
		moved = true
	case pads.Left.IsPressed(key):
		shadow.X = player.X
		player.X -= 1
		moved = true
	case pads.Right.IsPressed(key):
		shadow.X = player.X
		player.X += 1
		moved = true
	}
}

func draw() {
	if moved {
		shadow.Draw()
		player.Draw()
		moved = false
	}
}
