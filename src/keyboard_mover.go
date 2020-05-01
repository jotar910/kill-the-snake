package main

import (
	"math"

	"github.com/jotar910/kill-the-snake/src/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardMover struct {
	speed        float64
	spriteDrawer *SpriteDrawer

	container *Element
}

func NewKeyboardMover(container *Element, speed float64) (*KeyboardMover, error) {
	return &KeyboardMover{
		container:    container,
		spriteDrawer: container.getComponent(&SpriteDrawer{}).(*SpriteDrawer),
		speed:        speed,
	}, nil
}

func (km *KeyboardMover) onDraw(r *sdl.Renderer) error {
	return nil
}

func (km *KeyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	angle, ok := utils.AngleFromKeys(keys)
	if !ok {
		return nil
	}

	pos := km.container.Position
	newX := pos.X + (math.Cos(angle*math.Pi/180))*km.speed
	newY := pos.Y + (math.Sin(angle*math.Pi/180))*km.speed
	drawer := km.spriteDrawer
	limitX := float64(ScreenWidth - drawer.W)
	limitY := float64(ScreenHeight - drawer.H)

	if newX < 0 || newX > limitX {
		newX = pos.X
	}

	if newY < 0 || newY > limitY {
		newY = pos.Y
	}

	km.container.Position = Vector{X: newX, Y: newY}
	km.container.Angle = angle

	return nil
}

func (*KeyboardMover) onCollision(_ *Element) error {
	return nil
}
