package timing

import "device/gba"

func WaitForVBlank() {
	for gba.DISP.DISPSTAT.Get()&(1<<gba.DISPSTAT_VBLANK_Pos) == 0 {
	}
}
