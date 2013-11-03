package main

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/nsf/termbox-go"
	"time"
)

func drawAll() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func drawSnake() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	nodes := snake.Pos
	for e := nodes.Front(); e != nil; e = e.Next() {
		pos := e.Value.([2]int)
		termbox.SetCell(pos[0], pos[1], 'O', termbox.ColorYellow, termbox.ColorDefault)
	}
	termbox.Flush()
}

var snake *Snake

func main() {
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
				snake.MoveDown(false)
			case termbox.KeyArrowRight:
				snake.MoveRight(false)
			case termbox.KeyArrowUp:
				snake.MoveUp(false)
			case termbox.KeyArrowLeft:
				snake.MoveLeft(false)
			case termbox.KeyEsc:
				break loop
			}
		default:
			if snake.KeepGoing != nil {
				snake.KeepGoing(false)
			}
			drawSnake()
			time.Sleep(100 * time.Millisecond)
		}
	}
}
