package UI

import (
	"github.com/ITR13/turingMachine/frameEngine"
)

type Menu struct {
	screen   *frames.Screen
	cursor   *frames.Sprite
	elements []*MenuElement
}

type MenuElement struct {
	sprite    *frames.Sprite
	nextField *frames.Field
}

func MakeMenu(backgroundPath, cursorPath string,
	fe *frames.FrameEngine) (*Menu, error) {
	screen := fe.NewScreen()
	err := screen.SetBackground(backgroundPath)
	if err != nil {
		return nil, err
	}
	if cursorPath != "" {
		cursorTex, err := screen.GetTexture(cursorPath)
		cursor := screen.GetSpriteFromTexture(cursorTex)
		if err != nil {
			screen.Destroy()
			return nil, err
		}
		return &Menu{screen, cursor, make([]*MenuElement, 0)}, nil
	}
	return &Menu{screen, nil, make([]*MenuElement, 0)}, nil
}

func (menu *Menu) AddElement(x, y int, text string, next *frames.UpdateFunc) {

}
