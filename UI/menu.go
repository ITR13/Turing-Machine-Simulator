package UI

import "github.com/ITR13/turingMachine/graphics"

type Menu struct {
	screen   *graphics.Screen
	cursor   *graphics.Sprite
	elements []*MenuElement
}

type MenuElement struct {
	sprite *graphics.Screen
}
