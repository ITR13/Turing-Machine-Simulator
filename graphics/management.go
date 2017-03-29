package graphics

func (screen *Screen) SetBackground(path string) error {
	var texture *Texture
	if path != "" {
		texture, err := screen.GetTexture(path)
		if err != nil {
			return err
		}
	}

	if screen.background != nil {
		screen.background.Destroy()
	}
	screen.background = texture

	return nil
}
