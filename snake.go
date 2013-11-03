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

const (
	SnakeStateStop  = "stop"
	SnakeStateDown  = "down"
	SnakeStateUp    = "up"
	SnakeStateLeft  = "left"
	SnakeStateRight = "right"
)

func NewSnake() *Snake {
	s := &Snake{}

	pos := list.New()
	pos.PushFront([2]int{10, 10})
	s.Pos = pos

	movements := make(map[string]Movement)
	movements[SnakeStateStop] = [2]int{0, 0}
	movements[SnakeStateDown] = [2]int{0, 1}
	movements[SnakeStateUp] = [2]int{0, -1}
	movements[SnakeStateLeft] = [2]int{-1, 0}
	movements[SnakeStateRight] = [2]int{1, 0}
	s.Movements = movements

	s.State = SnakeStateStop

	return s
}

func (this *Snake) KeepGoing(eaten bool) {
	this.move(this.Movements[this.State])
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) MoveDown(eaten bool) {
	if this.isStateIn(SnakeStateUp, SnakeStateDown) {
		return
	}
	this.State = SnakeStateDown
	this.KeepGoing(eaten)
}

func (this *Snake) MoveUp(eaten bool) {
	if this.isStateIn(SnakeStateUp, SnakeStateDown) {
		return
	}
	this.State = SnakeStateUp
	this.KeepGoing(eaten)
}

func (this *Snake) MoveLeft(eaten bool) {
	if this.isStateIn(SnakeStateLeft, SnakeStateRight) {
		return
	}
	this.State = SnakeStateLeft
	this.KeepGoing(eaten)
}

func (this *Snake) MoveRight(eaten bool) {
	if this.isStateIn(SnakeStateLeft, SnakeStateRight) {
		return
	}
	this.State = SnakeStateRight
	this.KeepGoing(eaten)
}

func (this *Snake) isStateIn(states ...string) bool {
	var flag bool
	flag = false
	for _, state := range states {
		if this.State == state {
			flag = true
			break
		}
	}
	return flag
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
