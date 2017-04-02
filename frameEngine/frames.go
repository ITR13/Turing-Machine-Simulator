package frames

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
	"github.com/go-errors/errors"
)

func (fe *FrameEngine) Run(origin *Field) error {
	if fe.running {
		return errors.Errorf("Cannot run an already running FrameEngine")
	}
	fe.running = true
	defer func() {
		fe.running = false
	}()

	err := fe.initFE(origin)
	if err != nil {
		return err
	}
	for fe.currentField != nil {
		fe.update()
		if fe.error != nil {
			return errors.Errorf("Uncaught error: %v", fe.error)
		}
	}

	return nil
}

func (fe *FrameEngine) update() {
	fe.canPushOrPop = true
	cf := fe.currentField
	cf.screen.Render()
	cf.update()
	fe.canPushOrPop = false
}

func (fe *FrameEngine) PushField(field *Field) error {
	if !fe.canPushOrPop {
		fe.error = errors.Errorf("Tried pushing after having " +
			"pushed or popped or switched")
		return fe.error
	}
	if field == nil {
		fe.error = errors.Errorf("Tried pushing nil")
		return fe.error
	}
	if field.fe != fe {
		fe.error = errors.Errorf("Tried pushing field of other FrameEngine")
		return fe.error
	}

	fe.canPushOrPop = false
	fe.fields = append(fe.fields, fe.currentField)
	fe.currentField = field
	return nil
}

func (fe *FrameEngine) PopField() error {
	if !fe.canPushOrPop {
		fe.error = errors.Errorf("Tried popping after having " +
			"pushed or popped or switched")
		return fe.error
	}
	fe.currentField = nil
	l := len(fe.fields)
	if l > 0 {
		fe.currentField = fe.fields[l-1]
		fe.fields = fe.fields[:l-1]
	}
	return nil
}

func (screen *Screen) MakeField(update UpdateFunc) *Field {
	return &Field{screen, update, screen.fe}
}
