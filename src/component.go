package main

import "github.com/veandco/go-sdl2/sdl"

type Component interface {
	onDraw(r *sdl.Renderer) error
	onUpdate() error
	onCollision(*Element) error
}
