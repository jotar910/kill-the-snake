package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type StraightMover struct {
	speed        float64
	spriteDrawer *SpriteDrawer
	container    *Element
}

func NewStraightMover(container *Element, speed float64) (*StraightMover, error) {
	return &StraightMover{
		container:    container,
		spriteDrawer: container.getComponent(&SpriteDrawer{}).(*SpriteDrawer),
		speed:        speed,
	}, nil
}

func (sm *StraightMover) onDraw(r *sdl.Renderer) error {
	return nil
}

func (sm *StraightMover) onUpdate() error {
	pos := sm.container.Position
	angle := sm.container.Angle * math.Pi / 180.0

	newX := pos.X + math.Cos(angle)*sm.speed
	newY := pos.Y + math.Sin(angle)*sm.speed
	sm.container.Position = Vector{X: newX, Y: newY}

	drawer := sm.spriteDrawer
	limitX := float64(drawer.W + ScreenWidth)
	limitY := float64(drawer.H + ScreenHeight)

	if newX < -float64(drawer.W) || newX > limitX || newY < -float64(drawer.H) || newY > limitY {
		sm.container.Active = false
	}
	return nil
}
