package frames

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
		fe.canPushOrPop = false
		fe.update()
		if fe.error != nil {
			return errors.Errorf("Uncaught error: %v", fe.error)
		}
	}

	return nil
}

func (fe *FrameEngine) update() {
	fe.canPushOrPop = true
	fe.currentField.screen.Render()
	(*fe.currentField.update)()
	fe.canPushOrPop = false
}

func (fe *FrameEngine) PushField(field *Field) error {
	if !fe.canPushOrPop {
		fe.error = errors.Errorf("Tried pushing after having " +
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
