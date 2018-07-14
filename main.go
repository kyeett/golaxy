package main

import (
	"math"
	"sort"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var astroids []*Astroid
var gridSize float64 = 150

func main() {
	// Create initial astroids
	var padding float64 = 50
	for j := float64(0); j < 4; j++ {
		for i := float64(0); i < 3; i++ {
			astroids = append(astroids, NewRandomAstroid(pixel.V(padding-j*gridSize, padding+i*gridSize)))
		}
	}

	// Create scoreboard
	s := newScoreboard()

	// Start function that updates objects
	go func(scoreboard *scoreboard) {
		ticker := time.NewTicker(10 * time.Millisecond)
		for {
			<-ticker.C

			for _, a := range astroids {
				a.UpdatePosition()

				// Wrap astroids around board
				if a.pos.X > 450 {

					// Wrap astroid position
					a.pos.X = math.Mod(a.pos.X, 450) - 150

					// Update scoreboard
					scoreboard.counters["# astroids collected"]++
					scoreboard.counters["$ earned"] += a.value
				}
			}
		}
	}(s)

	pixelgl.Run(func() {
		run(s)
	})
}

func run(scoreboard *scoreboard) {

	cfg := pixelgl.WindowConfig{
		Title:  "Golaxy Rocks!",
		Bounds: pixel.R(0, 0, 450, 450+150),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	win.SetPos(pixel.V(2200, 100))
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(nil)
	imd.Color = colornames.White
	imd.EndShape = imdraw.RoundEndShape

	for !win.Closed() {
		imd.Clear()

		scoreboard.Update()

		for _, a := range astroids {

			imd.Push(a.corners...)
			imd.SetMatrix(pixel.IM.Rotated(pixel.V(25, 25), (math.Pi/180)*float64(a.angle)).Moved(a.pos))
			imd.Polygon(1)
		}

		// Clear and draw
		win.Clear(colornames.Black)
		scoreboard.Draw(win, pixel.IM)
		imd.Draw(win)
		win.Update()
	}
}

func sortedKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
