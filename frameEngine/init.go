package frames

import (
	"fmt"

	"github.com/ITR13/turingMachine/graphics"
)

type FrameEngine struct {
	fields       []*Field
	currentField *Field
	canPushOrPop bool
	error        error
}

type Field struct {
	screen *graphics.Screen
	update func()
}

func (fe *FrameEngine) Run(origin *Field) error {
	err := graphics.Init()
	if err != nil {
		return err
	}
	fe.currentField = origin
	fe.fields = make([]*Field, 0)
	fe.canPushOrPop = false
	for fe.currentField != nil {
		fe.update()
		if fe.error != nil {
			return fmt.Errorf("Uncaught error: %v", fe.error)
		}
	}

	return nil
}

func (fe *FrameEngine) update() {
	fe.canPushOrPop = true
	fe.currentField.screen.Render()
	fe.currentField.update()
	fe.canPushOrPop = false
}

func (fe *FrameEngine) PushField(field *Field) error {
	if !fe.canPushOrPop {
		fe.error = fmt.Errorf("Tried pushing after having " +
			"pushed or popped or switched")
		return fe.error
	}
	fe.canPushOrPop = false
	fe.fields = append(fe.fields, fe.currentField)
	fe.currentField = field
	return nil
}

func (fe *FrameEngine) PopField() error {
	if !fe.canPushOrPop {
		fe.error = fmt.Errorf("Tried popping after having " +
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
