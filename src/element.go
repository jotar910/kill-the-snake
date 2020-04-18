package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector struct {
	X, Y float64
}

func (v *Vector) String() string {
	return fmt.Sprintf("x: %.2f, y: %.2f", v.X, v.Y)
}

type Element struct {
	Position   Vector
	Rotation   float64
	Angle      float64
	Active     bool
	components []Component
}

func (e *Element) addComponent(component Component) {
	typ := reflect.TypeOf(component)
	for _, c := range e.components {
		if reflect.TypeOf(c) == typ {
			panic(errors.New("adding component to an element that already has the same component type"))
		}
	}
	e.components = append(e.components, component)
}

func (e *Element) getComponent(componentWithType Component) Component {
	typ := reflect.TypeOf(componentWithType)
	for _, c := range e.components {
		if reflect.TypeOf(c) == typ {
			return c
		}
	}
	panic(errors.New("component does not exist on element"))
}

func (e *Element) Draw(r *sdl.Renderer) error {
	for _, c := range e.components {
		err := c.onDraw(r)
		if err != nil {
			return errors.New(fmt.Sprintln("drawing component on element", err))
		}
	}
	return nil
}

func (e *Element) Update() error {
	for _, c := range e.components {
		err := c.onUpdate()
		if err != nil {
			return errors.New(fmt.Sprintln("updating component on element", err))
		}
	}
	return nil
}
