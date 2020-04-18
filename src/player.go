package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerFilename = "./assets/player/player_front_1.png"
	playerSpeed    = 0.05
	playerColdown  = 200 * time.Millisecond
)

func NewPlayer(r *sdl.Renderer) (*Element, error) {
	player := &Element{}

	player.Active = true
	player.Position = Vector{X: 0, Y: 0}
	player.Rotation = 0

	drawer, err := NewSpriteDrawer(r, playerFilename, player)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("creating sprite drawer component for player element", err))
	}
	player.addComponent(drawer)

	mover, err := NewKeyboardMover(player, playerSpeed)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("creating keyboard mover component for player element", err))
	}
	player.addComponent(mover)

	shooter, err := NewKeyboardShooter(player, playerColdown)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("creating keyboard shooter component for player element", err))
	}
	player.addComponent(shooter)

	return player, nil
}
