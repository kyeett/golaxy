package main

import (
	"math/rand"

	"github.com/faiface/pixel"
)

// Astroid represents a polygon with a current angle and position
type Astroid struct {
	corners    []pixel.Vec
	pos        pixel.Vec
	angle      int
	angleSpeed int
	value      int
}

// UpdatePosition incrementally updates astroid position
func (a *Astroid) UpdatePosition() {
	a.angle += a.angleSpeed
	a.pos.X += 2
}

// NewRandomAstroid creates an astroid with a random starting angle, angleSpeed and value
// Initial position is specified by user
func NewRandomAstroid(p pixel.Vec) *Astroid {
	angleSpeeds := []int{
		-2,
		-1,
		1,
		2,
	}
	return &Astroid{
		corners: []pixel.Vec{
			pixel.V(0, 0),
			pixel.V(25, 0),
			pixel.V(50, 25),
			pixel.V(25, 50),
			pixel.V(0, 50),
			pixel.V(0, 0),
		},
		pos:        p,
		angle:      rand.Intn(360),
		angleSpeed: angleSpeeds[rand.Intn(len(angleSpeeds))],
		value:      rand.Intn(100),
	}
}
