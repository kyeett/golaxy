package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type scoreboard struct {
	counters map[string]int
	*text.Text
}

func (s scoreboard) Update() {
	s.Clear()
	fmt.Fprintln(s, "Economics!")

	for _, k := range sortedKeys(s.counters) {
		fmt.Fprintf(s, "%v: %v\n", k, s.counters[k])
	}
}

func newScoreboard() *scoreboard {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	s := &scoreboard{
		make(map[string]int),
		text.New(pixel.V(50, 550), basicAtlas),
	}
	s.counters["# astroids collected"] = 0
	s.counters["$ earned"] = 0

	return s
}
