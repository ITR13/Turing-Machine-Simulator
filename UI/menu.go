package UI

/*
   This file is part of Turing Machine.

   Turing Machine is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   Turing Machine is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with Turing Machine.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"github.com/ITR13/turingMachine/frameEngine"
)

type Menu struct {
	field    *frames.Field
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
		return &Menu{nil, screen, cursor, make([]*MenuElement, 0)}, nil
	}
	return &Menu{nil, screen, nil, make([]*MenuElement, 0)}, nil
}

func (menu *Menu) AddElement(x, y int, text string, next *frames.Field) {

}

func (menu *Menu) GetField() *frames.Field {
	if menu.field == nil {
		current := 0
		menu.field = menu.screen.MakeField(func() {

		})
	}
	return menu.field
}
