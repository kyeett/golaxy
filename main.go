package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var astroids []*astroid
var gridSize float64 = 150
var counters map[string]int

func main() {

	counters = make(map[string]int)
	counters["# astroids collected"] = 0
	counters["$ earned"] = 0

	// Create astroids
	var padding float64 = 50
	for j := float64(0); j < 4; j++ {
		for i := float64(0); i < 3; i++ {
			astroids = append(astroids, newAstroid(pixel.V(padding-j*gridSize, padding+i*gridSize)))
		}
	}

	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		for {
			<-ticker.C

			for _, a := range astroids {
				a.angle += a.angleSpeed
				a.pos.X += 2
				if a.pos.X > 450 {
					a.pos.X = math.Mod(a.pos.X, 450) - 150
					counters["# astroids collected"]++
					counters["$ earned"] += a.value
				}
			}
		}
	}()

	pixelgl.Run(run)
}

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 450, 450+150),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	win.SetPos(pixel.V(2200, 100))
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Black
	//imd.SetMatrix(pixel.IM.Rotated(pixel.V(1, 0), 45))

	//imd.SetMatrix(pixel.IM.Rotated(pixel.V(1, 0), 45))

	for !win.Closed() {

		imd.Clear()
		imd.Color = colornames.White
		basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
		basicTxt := text.New(pixel.V(50, 550), basicAtlas)
		fmt.Fprintln(basicTxt, "Economics!")

		for _, k := range sortedKeys(counters) {
			fmt.Fprintf(basicTxt, "%v: %v\n", k, counters[k])
		}

		for _, a := range astroids {
			imd.EndShape = imdraw.RoundEndShape

			imd.Push(a.corners...)
			imd.SetMatrix(pixel.IM.Rotated(pixel.V(25, 25), (math.Pi/180)*float64(a.angle)).Moved(a.pos))
			imd.Polygon(1)
		}

		//Draw grid
		/*imd.SetMatrix(pixel.IM)
		for x := gridSize; x < cfg.Bounds.W(); x += gridSize {
			imd.Push(pixel.V(x, 0), pixel.V(x, cfg.Bounds.H()))
			imd.Line(1)
		}
		for y := gridSize; y < cfg.Bounds.H(); y += gridSize {
			imd.Push(pixel.V(0, y), pixel.V(cfg.Bounds.W(), y))
			imd.Line(1)
		}*/
		win.Clear(colornames.Black)
		basicTxt.Draw(win, pixel.IM)
		imd.Draw(win)
		win.Update()
	}
}

func newAstroid(p pixel.Vec) *astroid {
	angleSpeeds := []int{
		-2,
		-1,
		1,
		2,
	}
	return &astroid{
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

func sortedKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

type astroid struct {
	corners    []pixel.Vec
	pos        pixel.Vec
	angle      int
	angleSpeed int
	value      int
}
