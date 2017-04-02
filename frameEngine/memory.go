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
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

func (fe *FrameEngine) NewScreen() *Screen {
	g := fe.graphics
	mem := &Screen{len(g.allGraphics), nil,
		make([]*Sprite, 0), make([]*Texture, 0), fe}
	g.allGraphics = append(g.allGraphics, mem)
	return mem
}

func (mem *Screen) Destroy() {
	g := mem.fe.graphics
	if mem.index == -1 {
		panic("Handling destroyed screen")
	}
	mem.index = -1
	g.allGraphics = g.allGraphics[mem.index : mem.index+1]
	for i := mem.index; i < len(g.allGraphics); i++ {
		g.allGraphics[i].decrIndex()
	}
	mem.destroy()
}

func (mem *Screen) destroy() {
	if mem.background != nil {
		mem.background.destroy()
		mem.background = nil
	}
	for i := range mem.sprites {
		mem.sprites[i].texture = nil
	}
	for i := range mem.textures {
		mem.textures[i].destroy()
	}
	mem.sprites, mem.textures = make([]*Sprite, 0), make([]*Texture, 0)
}

func (tex *Texture) destroy() {
	if tex.texture == nil {
		panic("Handling destroyed texture")
	}
	tex.texture.Destroy()
	tex.texture = nil
	tex.src, tex.dst = nil, nil
}

func (screen *Screen) decrIndex() {
	screen.index--
	for i := range screen.textures {
		screen.textures[i].screenIndex--
	}
}

func (screen *Screen) GetSpriteFromTexture(texture *Texture) *Sprite {
	if texture == nil {
		panic("Tried to make sprite without texture")
	}
	if texture.screenIndex != screen.index {
		panic("Textures are screen-specific")
	}
	sprite := &Sprite{0, 0, false, texture}
	screen.sprites = append(screen.sprites, sprite)
	return sprite
}

func (screen *Screen) GetTexture(path string) (*Texture, error) {
	texture, err := screen.getTexture(path)
	if err != nil {
		return nil, err
	}
	screen.textures = append(screen.textures, texture)
	return texture, nil
}

func (screen *Screen) getTexture(path string) (*Texture, error) {
	surface, err := img.Load(path)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	tex, err := screen.fe.graphics.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}
	w, h := surface.W, surface.H

	return &Texture{
		screen.index,
		tex,
		&sdl.Rect{0, 0, w, h},
		&sdl.Rect{0, 0, w, h},
		screen.fe,
	}, nil
}
