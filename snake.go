package main

import (
	"container/list"
)

type Movement [2]int

type Snake struct {
	Pos       *list.List
	State     string
	Movements map[string]Movement
}

func NewSnake() *Snake {
	s := &Snake{}

	pos := list.New()
	pos.PushFront([2]int{10, 10})
	s.Pos = pos

	movements := make(map[string]Movement)
	movements["stop"] = [2]int{0, 0}
	movements["down"] = [2]int{0, 1}
	movements["up"] = [2]int{0, -1}
	movements["left"] = [2]int{-1, 0}
	movements["right"] = [2]int{1, 0}
	s.Movements = movements

	s.State = "stop"

	return s
}

func (this *Snake) KeepGoing(eaten bool) {
	this.move(this.Movements[this.State])
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) MoveDown(eaten bool) {
	if this.State == "up" {
		return
	}
	this.State = "down"
	this.KeepGoing(eaten)
}

func (this *Snake) MoveUp(eaten bool) {
	if this.State == "down" {
		return
	}
	this.State = "up"
	this.KeepGoing(eaten)
}

func (this *Snake) MoveLeft(eaten bool) {
	if this.State == "right" {
		return
	}
	this.State = "left"
	this.KeepGoing(eaten)
}

func (this *Snake) MoveRight(eaten bool) {
	if this.State == "left" {
		return
	}
	this.State = "right"
	this.KeepGoing(eaten)
}

func (this *Snake) move(movement [2]int) {
	headNode, headPos := this.addHead()
	headNode.Value = [2]int{headPos[0] + movement[0], headPos[1] + movement[1]}
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
