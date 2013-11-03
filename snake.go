package main

import (
	"container/list"
)

type Snake struct {
	Pos       *list.List
	KeepGoing func(bool)
}

func NewSnake() *Snake {
	pos := list.New()
	pos.PushFront([2]int{10, 10})
	return &Snake{
		Pos: pos,
	}
}

func (this *Snake) MoveDown(eaten bool) {
	this.move(0, 1)
	this.KeepGoing = this.MoveDown
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) MoveLeft(eaten bool) {
	this.move(-1, 0)
	this.KeepGoing = this.MoveLeft
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) MoveUp(eaten bool) {
	this.move(0, -1)
	this.KeepGoing = this.MoveUp
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) MoveRight(eaten bool) {
	this.move(1, 0)
	this.KeepGoing = this.MoveRight
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) move(x, y int) {
	headNode, headPos := this.addHead()
	headNode.Value = [2]int{headPos[0] + x, headPos[1] + y}
}

func (this *Snake) addHead() (headNode *list.Element, headPos [2]int) {
	node := this.Pos.Front()
	headNode = this.Pos.PushFront(node.Value)
	return headNode, headNode.Value.([2]int)
}

func (this *Snake) cutTail() {
	tail := this.Pos.Back()
	this.Pos.Remove(tail)
}
