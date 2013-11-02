package main

import (
	"container/list"
)

type Snake struct {
	Pos *list.List
}

func NewSnake() *Snake {
	pos := list.New()
	pos.PushFront([2]int{10, 10})
	return &Snake{
		Pos: pos,
	}
}

func (this *Snake) MoveDown(eaten bool) {
	headNode, headPos := this.addHead()
	headNode.Value = [2]int{headPos[0], headPos[1] + 1}
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) MoveLeft(eaten bool) {
	headNode, headPos := this.addHead()
	headNode.Value = [2]int{headPos[0] - 1, headPos[1]}
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) MoveUp(eaten bool) {
	headNode, headPos := this.addHead()
	headNode.Value = [2]int{headPos[0], headPos[1] - 1}
	if !eaten {
		this.cutTail()
	}
}

func (this *Snake) MoveRight(eaten bool) {
	headNode, headPos := this.addHead()
	headNode.Value = [2]int{headPos[0] + 1, headPos[1]}
	if !eaten {
		this.cutTail()
	}
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
