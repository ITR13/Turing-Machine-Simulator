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

func (screen *Screen) Render() {
	renderer := screen.fe.graphics.renderer
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
	tex.fe.graphics.renderer.Copy(tex.texture, tex.src, tex.dst)
}
