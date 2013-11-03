package main

import (
	"math/rand"
	"time"
)

type Scene struct {
	snake       *Snake
	Gem         [2]int
	Width       int
	Height      int
	Margin      int
	InnerWidth  int
	InnerHeight int
	InnerLeft   int
	InnerRight  int
	InnerTop    int
	InnerDown   int
}

func NewScene(snake *Snake, width, height, margin int) *Scene {
	rand.Seed(time.Now().UTC().UnixNano())
	s := &Scene{
		snake:  snake,
		Gem:    [2]int{-1, -1},
		Width:  width,
		Height: height,
		Margin: margin,
	}
	s.init()
	return s
}

func (this *Scene) init() {
	this.InnerTop = this.Margin
	this.InnerDown = this.Height - this.Margin
	this.InnerHeight = this.InnerDown - this.InnerTop - 1

	this.InnerLeft = this.Margin
	this.InnerRight = this.Width - this.Margin
	if (this.InnerRight-this.InnerLeft-1)%2 == 0 {
		this.InnerWidth = this.InnerRight - this.InnerLeft - 1
	} else {
		this.InnerRight = this.InnerRight + 1
		this.InnerWidth = this.InnerRight - this.InnerLeft - 1
	}

	this.snake.SetHead(this.InnerRight-2, this.InnerTop+1)
}

func (this *Scene) generateGem() {

  // x: InnerLeft+1 ~ InnerRight-2
  widthRange := this.InnerWidth / 2

	this.Gem = [2]int{
		this.InnerLeft + 2*(1 + rand.Int()%widthRange) - 1,
		1 + this.InnerTop + rand.Int()%this.InnerHeight,
	}
}

func (this *Scene) isSnakeOnGem() bool {
	_, headPos := this.snake.Head()
	return this.Gem == headPos
}

func (this *Scene) BounderCheck() {

	_, headPos := this.snake.Head()

	if headPos[0] >= this.InnerRight {
		this.snake.SetHead(headPos[0]-this.InnerWidth, headPos[1])
	} else if headPos[0] <= this.InnerLeft {
		this.snake.SetHead(headPos[0]+this.InnerWidth, headPos[1])
	}

	if headPos[1] >= this.InnerDown {
		this.snake.SetHead(headPos[0], headPos[1]-this.InnerHeight)
	} else if headPos[1] <= this.InnerTop {
		this.snake.SetHead(headPos[0], headPos[1]+this.InnerHeight)
	}

}
