package renderer

import (
	"battleground/constants"
	"battleground/state"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Go-SDL2 Events"
var winWidth, winHeight int32 = 800, 800

var window *sdl.Window
var renderer *sdl.Renderer
var event sdl.Event

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

func Render(world *state.World) bool {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return false

		}
	}
	renderer.Clear()
	for x := 0; x < world.Map.GetXn(); x++ {
		for y := 0; y < world.Map.GetYn(); y++ {
			block := &sdl.Rect{
				X: int32(constants.MapTileSize * x),
				Y: int32(constants.MapTileSize * y),
				W: constants.MapTileSize,
				H: constants.MapTileSize,
			}

			switch world.Map.GetTilesAt(x, y) {
			case 1:
				renderer.SetDrawColor(0, 0, 255, 255)
				break
			case 0:
				renderer.SetDrawColor(255, 255, 255, 255)
				break
			case 2:
				renderer.SetDrawColor(0, 0, 0, 255)
				break

			}

			renderer.FillRect(block)
		}
	}

	renderer.SetDrawColor(0, 0, 0, 255)

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

	// renderer.SetDrawColor(255, 255, 255, 255)
	// renderer.FillRect(&sdl.Rect{
	// 	X: 900,
	// 	Y: 900,
	// 	W: 20,
	// 	H: 20,
	// })

	renderer.Present()
	return true
	// renderer.SetDrawColor(255, 0, 0, 255)
	// renderer.Clear()

}
