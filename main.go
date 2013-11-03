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
	termbox.Flush()
}

func drawSnake() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	nodes := snake.Nodes
	for e := nodes.Front(); e != nil; e = e.Next() {
		pos := e.Value.([2]int)
		termbox.SetCell(pos[0], pos[1], '[', termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(pos[0]+1, pos[1], ']', termbox.ColorYellow, termbox.ColorDefault)
	}
	termbox.Flush()
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

func main() {

	logFile, _ := os.OpenFile("logs", os.O_RDWR|os.O_APPEND, 0660)
	defer logFile.Close()
	log.SetOutput(logFile)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	snake = NewSnake()
	drawSnake()

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
				snake.TurnDown(false)
			case termbox.KeyArrowRight:
				snake.TurnRight(false)
			case termbox.KeyArrowUp:
				snake.TurnUp(false)
			case termbox.KeyArrowLeft:
				snake.TurnLeft(false)
			case termbox.KeyEsc:
				break loop
			}
		default:
			snake.KeepGoing(false)
			checkOutOfBound()
			drawSnake()
			log.Println(snake.Nodes.Front().Value)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
