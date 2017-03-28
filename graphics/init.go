package graphics

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	screenWidth  int = 1280 / 2
	screenHeight int = 800 / 2
)

var (
	window   *sdl.Window
	renderer *sdl.Renderer
	font     *ttf.Font
)

type Screen struct {
	index      int
	background *sdl.Texture
	bsrc, bdst *sdl.Rect
	sprites    []*Sprite
	textures   []*Texture
}

type Sprite struct {
	X, Y    int
	Display bool
	texture *Texture
}

type Texture struct {
	screenIndex int
	texture     *sdl.Texture
	src, dst    *sdl.Rect
}

func Init() error {
	sdl.Init(sdl.INIT_EVERYTHING)

	runtime.LockOSThread()
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return err
	}

	window, err = sdl.CreateWindow("Turing Machine", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, int(screenWidth), int(screenHeight),
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE|sdl.RENDERER_PRESENTVSYNC)
	return err

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	renderer.Clear()

	err = ttf.Init()
	if err != nil {
		return err
	}
	font, err = ttf.OpenFont("./font/Play-Bold.ttf", 20)
	if err != nil {
		return err
	}
	allGraphics = make([]*Screen, 0)

	return nil
}

func CleanUp() {
	if window != nil {
		window.Destroy()
		window = nil
	}
	if renderer != nil {
		renderer.Destroy()
		renderer = nil
	}
	if font != nil {
		font.Close()
		font = nil
	}
	for i := range allGraphics {
		allGraphics[i].destroy()
		allGraphics = make([]*Screen, 0)
	}
}
