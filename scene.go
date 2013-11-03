package main

import (
	"math"
	"math/rand"
	"time"
)

type Scene struct {
	snake  *Snake
	Gem    [2]int
	Width  int
	Height int
}

func NewScene(snake *Snake, width, height int) *Scene {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Scene{
		snake:  snake,
		Gem:    [2]int{-1, -1},
		Width:  width,
		Height: height,
	}
}

func (this *Scene) generateGem() {
	this.Gem = [2]int{
		int(math.Mod(float64(rand.Int()), float64(this.Width))),
		int(math.Mod(float64(rand.Int()), float64(this.Height))),
	}
}

func (this *Scene) isSnakeOnGem() bool {
	_, headPos := this.snake.Head()
	return this.Gem == headPos
}
