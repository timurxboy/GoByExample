package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	loader := strings.ToUpper(t.talk())
	fmt.Println(loader)
}

type crater struct{}

type starship struct {
	laser
}

func main() {

	var t interface {
		talk() string
	}

	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(2))
	// shout(crater{})

	s := starship{laser(3)}

	fmt.Println(s.talk())
	shout(s)
}
