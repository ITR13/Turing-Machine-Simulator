package frames

func (screen *Screen) SetBackground(path string) error {
	var texture *Texture
	var err error
	if path != "" {
		texture, err = screen.getTexture(path)
		if err != nil {
			return err
		}
	}

	if screen.background != nil {
		screen.background.destroy()
	}
	screen.background = texture

	return nil
}
