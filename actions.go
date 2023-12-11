package main

import (
	"image"
	"math/rand"

	"github.com/go-vgo/robotgo"
)

func randomClick(area image.Rectangle) {
	x := area.Min.X + rand.Intn(area.Max.X-area.Min.X)
	y := area.Min.Y + rand.Intn(area.Max.Y-area.Min.Y)
	robotgo.MoveClick(int(x), int(y))
}
