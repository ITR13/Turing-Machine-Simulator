package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var allGraphics []*Screen

func NewScreen() *Screen {
	mem := &Screen{len(allGraphics), nil, nil, nil,
		make([]*Sprite, 0), make([]*Texture, 0)}
	allGraphics = append(allGraphics, mem)
	return mem
}

func (mem *Screen) Destroy() {
	if mem.index == -1 {
		panic("Handling destroyed screen")
	}
	mem.index = -1
	allGraphics = allGraphics[mem.index : mem.index+1]
	for i := mem.index; i < len(allGraphics); i++ {
		allGraphics[i].decrIndex()
	}
	mem.destroy()
}

func (mem *Screen) destroy() {
	mem.background.Destroy()
	mem.background = nil
	mem.bsrc, mem.bdst = nil, nil
	for i := range mem.sprites {
		mem.sprites[i].texture = nil
	}
	for i := range mem.textures {
		mem.textures[i].destroy()
	}
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

func (screen *Screen) GetSpriteFromTexture(texture *Texture) {
	if texture == nil {
		panic("Tried to make sprite without texture")
	}
	if texture.screenIndex != screen.index {
		panic("Textures are screen-specific")
	}
	screen.sprites = append(screen.sprites, &Sprite{
		0, 0, false, texture,
	})
}

func (screen *Screen) GetTexture(path string) (*Texture, error) {
	texture, err := getTexture(path)
	if err != nil {
		return err
	}
	screen.textures = append(screen.textures, texture)
}

func (screen *Screen) getTexture(path string) (*Texture, error) {
	surface, err := img.Load(path)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	tex, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}
	w, h := surface.W, surface.H

	return &Texture{
		screen.index,
		tex,
		&sdl.Rect{0, 0, w, h},
		&sdl.Rect{0, 0, w, h},
	}, nil
}
