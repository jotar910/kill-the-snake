package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ScreenWidth  = 760
	ScreenHeight = 600
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL: ", err)
		os.Exit(1)
	}

	screenWidth := ScreenWidth
	screenHeight := ScreenHeight

	window, err := sdl.CreateWindow(
		"Gaming in Go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(screenWidth), int32(screenHeight),
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window: ", err)
		os.Exit(1)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer: ", err)
		os.Exit(1)
	}
	defer renderer.Destroy()

	err = img.Init(img.INIT_JPG | img.INIT_PNG)
	if err != nil {
		fmt.Println("initializing image lib: ", err)
		os.Exit(1)
	}
	defer img.Quit()

	var elements []*Element

	bullets := InitBullets(renderer, 30)
	for _, bullet := range bullets {
		elements = append(elements, bullet)
	}
	playerInst, err := NewPlayer(renderer)
	if err != nil {
		fmt.Println("creating new player", err)
		os.Exit(1)
	}
	elements = append(elements, playerInst)
	snakes, err := NewSnakes(renderer, SNAKE_4, SNAKE_2)
	if err != nil {
		fmt.Println("creating new snakes", err)
		os.Exit(1)
	}
	for _, snake := range snakes {
		elements = append(elements, snake)
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 1)
		renderer.Clear()

		for _, element := range elements {
			if !element.Active {
				continue
			}
			err := element.Update()
			if err != nil {
				fmt.Println("update element", err)
				os.Exit(1)
			}
			err = element.Collision(elements)
			if err != nil {
				fmt.Println("update element", err)
				os.Exit(1)
			}
			err = element.Draw(renderer)
			if err != nil {
				fmt.Println("drawing element", err)
				os.Exit(1)
			}
		}

		renderer.Present()
	}
}
