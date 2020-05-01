package main

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const snakeFilename = "./assets/snakes/snake%d.png"

type SnakeType = int8

const (
	SNAKE_1 int8 = iota + 1
	SNAKE_2
	SNAKE_3
	SNAKE_4
)

func NewSnake(r *sdl.Renderer, typ SnakeType) (*Element, error) {
	snake := &Element{}

	drawer, err := NewSpriteDrawer(r, fmt.Sprintf(snakeFilename, typ), snake)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("creting sprite drawer component for snake element", err))
	}
	snake.addComponent(drawer)

	snake.Position.X = float64(ScreenWidth - drawer.W)
	snake.Position.Y = float64(ScreenHeight - drawer.H)
	snake.Active = true
	snake.Rotation = 0

	snake.BodyHit = Body{
		radius:   15,
		position: &snake.Position,
	}

	bulletCollision, err := NewBulletCollision(func(_ *Body) { snake.Active = false }, &snake.BodyHit)
	if err != nil {
		return nil, errors.New("creating the bullet collision for the snake")
	}
	snake.Collisions = append(snake.Collisions, bulletCollision)

	return snake, nil
}

func NewSnakes(r *sdl.Renderer, typs ...SnakeType) ([]*Element, error) {
	var snakes []*Element
	for i, typ := range typs {
		snake, err := NewSnake(r, typ)
		if err != nil {
			return nil, err
		}
		drawer := snake.getComponent(&SpriteDrawer{}).(*SpriteDrawer)
		snake.Position.X -= float64(int32(i) * drawer.W)
		snakes = append(snakes, snake)
	}
	return snakes, nil
}
