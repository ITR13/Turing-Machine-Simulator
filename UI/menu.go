package UI

import (
	"github.com/ITR13/turingMachine/frameEngine"
	"github.com/ITR13/turingMachine/graphics"
)

type Menu struct {
	screen   *graphics.Screen
	cursor   *graphics.Sprite
	elements []*MenuElement
}

type MenuElement struct {
	sprite    *graphics.Sprite
	nextField *frames.Field
}

func MakeMenu(backgroundPath, cursorPath string) (*Menu, err) {
	screen := graphics.NewScreen()
	err := screen.SetBacground(backgroundPath)
	if err != nil {
		return err
	}
	if cursorPath != "" {
		cursor, err := screen.GetTexture(cursorPath)
		if err != nil {
			screen.Destroy()
			return err
		}
		return &Menu{screen, cursor, make([]*MenuElement, 0)}, nil
	}
	return &Menu{screen, nil, make([]*MenuElement, 0)}, nil
}
