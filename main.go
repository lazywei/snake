package main

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"time"
)

func drawAll() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawGem()
	drawSnake()
	termbox.Flush()
}

func drawSnake() {
	head, headPos := snake.Head()

  termbox.SetCell(headPos[0], headPos[1], '(', termbox.ColorYellow, termbox.ColorDefault)
  termbox.SetCell(headPos[0]+1, headPos[1], ')', termbox.ColorYellow, termbox.ColorDefault)

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

func checkOutOfBound() {
	w, h := termbox.Size()
	_, headPos := snake.Head()

	if headPos[0] >= w {
		snake.SetHead(headPos[0]-w, headPos[1])
	} else if headPos[0] < 0 {
		snake.SetHead(headPos[0]+w, headPos[1])
	}

	if headPos[1] >= h {
		snake.SetHead(headPos[0], headPos[1]-h)
	} else if headPos[1] < 0 {
		snake.SetHead(headPos[0], headPos[1]+h)
	}

}

var snake *Snake
var scene *Scene

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
	scene = NewScene(snake, w, h)
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
			checkOutOfBound()
			drawAll()
			log.Println(snake.Nodes.Front().Value)
			time.Sleep(130 * time.Millisecond)
		}
	}
}
