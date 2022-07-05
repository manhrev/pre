package renderer

import (
	"battleground/constants"
	"battleground/state"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Go-SDL2 Events"
var winWidth, winHeight int32 = 800, 600

var window *sdl.Window
var renderer *sdl.Renderer

func Destroy() {
	sdl.Quit()
	window.Destroy()
	renderer.Destroy()
}

func Init() {

	var err error

	if e := sdl.Init(sdl.INIT_EVERYTHING); e != nil {
		panic(sdl.GetError())
	}

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
}

func Render(world *state.World) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	var rects []*sdl.Rect
	for _, obj := range world.Objects {
		rects = append(rects, &sdl.Rect{
			X: obj.Position().X - constants.PlayerSize,
			Y: obj.Position().Y - constants.PlayerSize,
			W: constants.PlayerSize * 2,
			H: constants.PlayerSize * 2})
	}

	for _, rect := range rects {
		renderer.SetDrawColor(0, 255, 255, 255)
		renderer.FillRect(rect)
	}

	renderer.Present()

	// renderer.SetDrawColor(255, 0, 0, 255)
	// renderer.Clear()

}
