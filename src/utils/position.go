package utils

import "github.com/veandco/go-sdl2/sdl"

func AngleFromKeys(keys []uint8) (float64, bool) {
	angleY := -1.0
	angleX := -1.0

	if keys[sdl.SCANCODE_LEFT] == 1 {
		angleX = 180.0
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		angleX = 360.0
	}

	if keys[sdl.SCANCODE_UP] == 1 {
		angleY = 270.0
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		angleY = 90.0
	}

	if angleX < 0 {
		angleX = angleY
	}
	if angleX == 360 && angleY == 90 {
		angleX = 0
	} else if angleY < 0 {
		angleY = angleX
	}
	if angleX < 0 || angleY < 0 {
		return -1, false
	}

	return (angleX + angleY) / 2.0, true
}
