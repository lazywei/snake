package main

import (
	"container/list"
)

type Movement [2]int

type Snake struct {
	Nodes     *list.List
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

	nodes := list.New()
	nodes.PushBack([2]int{10, 10})
	s.Nodes = nodes

	movements := make(map[string]Movement)
	movements[SnakeStateStop] = [2]int{0, 0}
	movements[SnakeStateDown] = [2]int{0, 1}
	movements[SnakeStateUp] = [2]int{0, -1}
	movements[SnakeStateLeft] = [2]int{-2, 0}
	movements[SnakeStateRight] = [2]int{2, 0}
	s.Movements = movements

	s.State = SnakeStateStop

	return s
}

func (this *Snake) SetHead(x, y int) {
	head := this.Nodes.Front()
	head.Value = [2]int{x, y}
}

func (this *Snake) Head() (headNode *list.Element, headPos [2]int) {
	headNode = this.Nodes.Front()
	return headNode, headNode.Value.([2]int)
}

func (this *Snake) KeepGoing(eaten bool) {
	this.move(this.Movements[this.State])
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) TurnDown() {
	if this.isStateIn(SnakeStateUp, SnakeStateDown) {
		return
	}
	this.State = SnakeStateDown
}

func (this *Snake) TurnUp() {
	if this.isStateIn(SnakeStateUp, SnakeStateDown) {
		return
	}
	this.State = SnakeStateUp
}

func (this *Snake) TurnLeft() {
	if this.isStateIn(SnakeStateLeft, SnakeStateRight) {
		return
	}
	this.State = SnakeStateLeft
}

func (this *Snake) TurnRight() {
	if this.isStateIn(SnakeStateLeft, SnakeStateRight) {
		return
	}
	this.State = SnakeStateRight
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
	this.addHead()
	headNode, headPos := this.Head()
	headNode.Value = [2]int{headPos[0] + movement[0], headPos[1] + movement[1]}
}

func (this *Snake) addHead() {
	node := this.Nodes.Front()
	this.Nodes.PushFront(node.Value)
}

func (this *Snake) cutTail() {
	tail := this.Nodes.Back()
	this.Nodes.Remove(tail)
}
