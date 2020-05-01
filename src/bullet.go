package main

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletFilename = "./assets/bullet/bullet.png"
	bulletSpeed    = 0.2
)

func InitBullets(r *sdl.Renderer, n int) []*Element {
	for i := n; i > 0; i-- {
		bullet, err := newBullet(r)
		if err != nil {
			panic(err)
		}
		bulletPool = append(bulletPool, bullet)
	}
	return bulletPool
}

func NewBulletFromPool() (*Element, bool) {
	for _, b := range bulletPool {
		if !b.Active {
			return b, true
		}
	}
	return nil, false
}

var bulletPool []*Element

func newBullet(r *sdl.Renderer) (*Element, error) {
	bullet := &Element{}

	bullet.Position = Vector{X: 0, Y: 0}
	bullet.Rotation = 0
	bullet.Active = false
	bullet.Tag = "bullet"

	drawer, err := NewSpriteDrawer(r, bulletFilename, bullet)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("creating sprite drawer component for bullet element", err))
	}
	bullet.addComponent(drawer)

	mover, err := NewBulletMover(bullet, bulletSpeed)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("creating straight mover component for bullet element", err))
	}
	bullet.addComponent(mover)

	hitter := Body{
		radius:   7,
		position: &bullet.Position,
	}
	bullet.Hitters = append(bullet.Hitters, hitter)

	return bullet, nil
}
