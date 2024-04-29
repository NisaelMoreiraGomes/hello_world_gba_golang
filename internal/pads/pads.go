package pads

import "device/gba"

type pad struct {
	pos int16
}

var (
	A = pad{gba.KEYINPUT_BUTTON_A_Pos}
	B = pad{gba.KEYINPUT_BUTTON_B_Pos}
	L = pad{gba.KEYINPUT_BUTTON_L_Pos}
	R = pad{gba.KEYINPUT_BUTTON_R_Pos}

	Up    = pad{gba.KEYINPUT_BUTTON_UP_Pos}
	Left  = pad{gba.KEYINPUT_BUTTON_LEFT_Pos}
	Right = pad{gba.KEYINPUT_BUTTON_RIGHT_Pos}
	Down  = pad{gba.KEYINPUT_BUTTON_DOWN_Pos}

	Select = pad{gba.KEYINPUT_BUTTON_SELECT_Pos}
	Start  = pad{gba.KEYINPUT_BUTTON_START_Pos}
)

const (
	padPressed  uint16 = 0
	padReleased uint16 = 1
)

func ReadPad() uint16 {
	return gba.KEY.INPUT.Get()
}

func (p pad) IsPressed(key uint16) bool {
	return key&(1<<p.pos) == padPressed
}

func (p pad) IsReleased(key uint16) bool {
	return key&(1<<p.pos) == padReleased
}
