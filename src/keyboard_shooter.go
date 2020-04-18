package main

import (
	"time"

	"github.com/jotar910/kill-the-snake/src/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardShooter struct {
	lastShoot    time.Time
	coldown      time.Duration
	spriteDrawer *SpriteDrawer
	container    *Element
}

func NewKeyboardShooter(container *Element, coldown time.Duration) (*KeyboardShooter, error) {
	return &KeyboardShooter{
		coldown:      coldown,
		container:    container,
		spriteDrawer: container.getComponent(&SpriteDrawer{}).(*SpriteDrawer),
	}, nil
}

func (ks *KeyboardShooter) onDraw(r *sdl.Renderer) error {
	return nil
}

func (ks *KeyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()
	drawer := ks.spriteDrawer
	pos := ks.container.Position

	angle, ok := utils.AngleFromKeys(keys)
	if !ok {
		return nil
	}

	if keys[sdl.SCANCODE_SPACE] == 1 && time.Since(ks.lastShoot) >= ks.coldown {
		ks.shoot(pos.X+float64(drawer.W/2.0), pos.Y+float64(drawer.H/2.0), angle)
	}
	return nil
}

func (ks *KeyboardShooter) shoot(x, y float64, angle float64) {
	if b, ok := NewBulletFromPool(); ok {
		b.Active = true
		b.Position.X = x
		b.Position.Y = y
		b.Rotation = ks.container.Rotation
		b.Angle = angle
		ks.lastShoot = time.Now()
	}
}
