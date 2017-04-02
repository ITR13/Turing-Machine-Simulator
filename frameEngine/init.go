package frames

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	screenWidth  int = 1280 / 2
	screenHeight int = 800 / 2
)

var hasInitSDL bool

type FrameEngine struct {
	running      bool
	fields       []*Field
	currentField *Field
	canPushOrPop bool
	error        error
	graphics     *Graphics //This really doesn't need to be a pointer, but
}

type Field struct {
	screen Screen
	update *UpdateFunc
}

type UpdateFunc func()

type Graphics struct {
	window      *sdl.Window
	renderer    *sdl.Renderer
	font        *ttf.Font
	allGraphics []*Screen
}

type Screen struct {
	index      int
	background *Texture
	sprites    []*Sprite
	textures   []*Texture
	fe         *FrameEngine
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
	fe          *FrameEngine
}

func (fe *FrameEngine) initFE(origin *Field) error {
	if fe.graphics == nil {
		fe.graphics = &Graphics{}
	} else {
		fe.graphics.CleanUp()
	}

	err := fe.graphics.initGraphics()
	if err != nil {
		return err
	}

	fe.currentField = origin
	fe.fields = make([]*Field, 0)
	return nil
}

func (g *Graphics) initGraphics() error {
	runtime.LockOSThread()
	var err error
	if !hasInitSDL {
		sdl.Init(sdl.INIT_EVERYTHING)
		err = sdl.Init(sdl.INIT_EVERYTHING)
		if err != nil {
			return err
		}
		err = ttf.Init()
		if err != nil {
			return err
		}
		hasInitSDL = true
	}

	g.window, err = sdl.CreateWindow("Turing Machine", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, int(screenWidth), int(screenHeight),
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return err
	}

	g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	g.renderer.Clear()

	g.font, err = ttf.OpenFont("./font/Play-Bold.ttf", 20)
	if err != nil {
		return err
	}
	g.allGraphics = make([]*Screen, 0)

	return nil
}

func (g *Graphics) CleanUp() {
	for i := range g.allGraphics {
		g.allGraphics[i].destroy()
	}
	g.allGraphics = make([]*Screen, 0)
	if g.font != nil {
		g.font.Close()
		g.font = nil
	}
	if g.window != nil {
		g.window.Destroy()
		g.window = nil
	}
	if g.renderer != nil {
		g.renderer.Destroy()
		g.renderer = nil
	}
}
