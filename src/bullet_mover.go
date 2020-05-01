package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type BulletMover struct {
	speed        float64
	spriteDrawer *SpriteDrawer
	container    *Element
}

func NewBulletMover(container *Element, speed float64) (*BulletMover, error) {
	return &BulletMover{
		container:    container,
		spriteDrawer: container.getComponent(&SpriteDrawer{}).(*SpriteDrawer),
		speed:        speed,
	}, nil
}

func (sm *BulletMover) onDraw(r *sdl.Renderer) error {
	return nil
}

func (sm *BulletMover) onUpdate() error {
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

func (bm *BulletMover) onCollision(e *Element) error {
	bm.container.Active = false
	return nil
}
