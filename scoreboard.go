package main

import (
	"fmt"
	"sort"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

type scoreboard struct {
	counters map[string]int
	*text.Text
}

func (s scoreboard) PrepareDraw() {
	s.Clear()
	fmt.Fprintln(s, "Economics!")

	for _, k := range sortedKeys(s.counters) {
		fmt.Fprintf(s, "%v: %v\n", k, s.counters[k])
	}
}

func newScoreboard() *scoreboard {
	basicAtlas := text.Atlas7x13

	s := &scoreboard{
		make(map[string]int),
		text.New(pixel.V(50, 550), basicAtlas),
	}
	s.counters["# astroids collected"] = 0
	s.counters["$ earned"] = 0
	s.counters["Time spent"] = 0

	return s
}

func sortedKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
