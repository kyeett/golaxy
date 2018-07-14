package main

import (
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
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
func NewRandomAstroid(p pixel.Vec) Astroid {
	angleSpeeds := []int{
		-2,
		-1,
		1,
		2,
	}
	return Astroid{
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

// Astroids is just an array of Astroid pointers
type Astroids []*Astroid

// PrepareDraw pushes the vertices of all astroids to the IMDraw object
// and applies tranformation to them
func (as Astroids) PrepareDraw(imd *imdraw.IMDraw) {
	for _, a := range as {
		imd.Push(a.corners...)
		imd.SetMatrix(pixel.IM.Rotated(pixel.V(25, 25), (math.Pi/180)*float64(a.angle)).Moved(a.pos))
		imd.Polygon(1)
	}
}
