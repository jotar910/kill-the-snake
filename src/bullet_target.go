package main

import "github.com/veandco/go-sdl2/sdl"

type BulletTarget struct {
	container *Element
}

func NewBulletTarget(e *Element) (*BulletTarget, error) {
	return &BulletTarget{e}, nil
}

func (bt *BulletTarget) onDraw(r *sdl.Renderer) error {
	return nil
}

func (bt *BulletTarget) onUpdate() error {
	return nil
}

func (bt *BulletTarget) onCollision(other *Element) error {
	if other.Tag == "bullet" {
		bt.container.Active = false
	}
	return nil
}
