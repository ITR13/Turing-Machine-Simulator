package frames

func (screen *Screen) Render() {
	renderer.SetRenderTarget(nil)
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	if screen.background != nil {
		renderer.Copy(screen.background.texture,
			screen.background.src, screen.background.dst)
	}
	for i := range screen.sprites {
		screen.sprites[i].render()
	}
}

func (sprite *Sprite) render() {
	if sprite.texture == nil {
		panic("Handling sprite without texture")
	}
	sprite.texture.renderAt(int32(sprite.X), int32(sprite.Y))
}

func (tex *Texture) renderAt(x, y int32) {
	tex.dst.X, tex.dst.Y = x, y
	renderer.Copy(tex.texture, tex.src, tex.dst)
}
