package main

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"time"
)

const (
  margin = 4
)

var (
  snake *Snake
  scene *Scene
)

func drawAll() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawGem()
	drawSnake()
  drawBoundary()
	termbox.Flush()
}

func drawSnake() {
	head, headPos := snake.Head()

	termbox.SetCell(headPos[0], headPos[1], ' ', termbox.ColorDefault, termbox.ColorYellow)
	termbox.SetCell(headPos[0]+1, headPos[1], ' ', termbox.ColorDefault, termbox.ColorYellow)

	for e := head.Next(); e != nil; e = e.Next() {
		pos := e.Value.([2]int)
		termbox.SetCell(pos[0], pos[1], '[', termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(pos[0]+1, pos[1], ']', termbox.ColorYellow, termbox.ColorDefault)
	}
}

func drawGem() {
	pos := scene.Gem
	termbox.SetCell(pos[0], pos[1], ' ', termbox.ColorDefault, termbox.ColorRed)
	termbox.SetCell(pos[0]+1, pos[1], ' ', termbox.ColorDefault, termbox.ColorRed)
}

func drawBoundary() {

	left := scene.InnerLeft
	right := scene.InnerRight

	top := scene.InnerTop
	down := scene.InnerDown

  for x := left; x <= right; x++ {
    termbox.SetCell(x, top, '-', termbox.ColorBlue, termbox.ColorDefault)
    termbox.SetCell(x, down, '-', termbox.ColorBlue, termbox.ColorDefault)
  }

  for y := top + 1; y < down; y++ {
    termbox.SetCell(left, y, '|', termbox.ColorBlue, termbox.ColorDefault)
    termbox.SetCell(right, y, '|', termbox.ColorBlue, termbox.ColorDefault)
  }
}

func main() {

	logFile, _ := os.OpenFile("logs", os.O_RDWR|os.O_APPEND, 0660)
	defer logFile.Close()
	log.SetOutput(logFile)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	w, h := termbox.Size()
	snake = NewSnake()
	scene = NewScene(snake, w, h, margin)
	scene.generateGem()

	drawAll()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

loop:
	for {
		select {
		case ev := <-eventQueue:
			switch ev.Key {
			case termbox.KeyArrowDown:
				snake.TurnDown()
			case termbox.KeyArrowRight:
				snake.TurnRight()
			case termbox.KeyArrowUp:
				snake.TurnUp()
			case termbox.KeyArrowLeft:
				snake.TurnLeft()
			case termbox.KeyEsc:
				break loop
			}
		default:
			if scene.isSnakeOnGem() {
				snake.KeepGoing(true)
				scene.generateGem()
			} else {
				snake.KeepGoing(false)
			}
			scene.BounderCheck()
			drawAll()
			log.Println(snake.Nodes.Front().Value)
			time.Sleep(150 * time.Millisecond)
		}
	}
}
