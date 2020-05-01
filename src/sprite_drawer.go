package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteDrawer struct {
	Tex  *sdl.Texture
	W, H int32

	container *Element
}

func NewSpriteDrawer(r *sdl.Renderer, filename string, container *Element) (*SpriteDrawer, error) {
	img, err := img.Load(filename)
	if err != nil {
		return nil, fmt.Errorf("loading %q png file: %v", filename, err)
	}
	defer img.Free()

	texture, err := r.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("creating texture:%v", err)
	}

	return &SpriteDrawer{Tex: texture, W: img.W, H: img.H, container: container}, nil
}

func (sd *SpriteDrawer) onDraw(r *sdl.Renderer) error {
	c := sd.container

	r.CopyEx(sd.Tex,
		&sdl.Rect{X: 0, Y: 0, W: sd.W, H: sd.H},
		&sdl.Rect{X: int32(c.Position.X), Y: int32(c.Position.Y), W: sd.W, H: sd.H},
		sd.container.Rotation,
		&sdl.Point{X: int32(sd.W / 2.0), Y: int32(sd.H / 2.0)},
		sdl.FLIP_NONE)

	return nil
}

func (sd *SpriteDrawer) onUpdate() error {
	return nil
}

func (*SpriteDrawer) onCollision(_ *Element) error {
	return nil
}
