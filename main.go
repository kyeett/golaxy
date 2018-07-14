package main

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var gridSize float64 = 150

type game struct {
	scoreboard *scoreboard
	astroids   Astroids
}

func (g game) UpdatePositions() {
	for _, a := range g.astroids {
		a.UpdatePosition()

		// Wrap astroids around board
		if a.pos.X > 450 {

			// Wrap astroid position
			a.pos.X = math.Mod(a.pos.X, 450) - 150

			// Update scoreboard
			g.scoreboard.counters["# astroids collected"]++
			g.scoreboard.counters["$ earned"] += a.value
		}
	}
	g.scoreboard.counters["Time spent"]++
}

func (g game) run() {

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Golaxy Rocks!",
		Bounds: pixel.R(0, 0, 450, 450+150),
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}

	// Used only when developing
	win.SetPos(pixel.V(2200, 100))

	imd := imdraw.New(nil)
	imd.Color = colornames.White
	imd.EndShape = imdraw.RoundEndShape

	ticker := time.NewTicker(10 * time.Millisecond)
	for !win.Closed() {
		// Limit the update frequency to once 100 times per second
		<-ticker.C

		// Updates positions and score
		g.UpdatePositions()

		// Clear main IMDraw
		imd.Clear()

		g.scoreboard.PrepareDraw()
		g.astroids.PrepareDraw(imd)

		// Clear and draw
		win.Clear(colornames.Black)
		g.scoreboard.Draw(win, pixel.IM)
		imd.Draw(win)
		win.Update()
	}
}

func main() {

	// Initialize game
	g := game{
		newScoreboard(),
		[]*Astroid{},
	}

	// Create astroids
	var padding float64 = 50
	for j := float64(0); j < 4; j++ {
		for i := float64(0); i < 3; i++ {
			a := NewRandomAstroid(pixel.V(padding-j*gridSize, padding+i*gridSize))
			g.astroids = append(g.astroids, &a)
		}
	}

	// Start main loop
	pixelgl.Run(g.run)
}
