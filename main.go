package main

import (
	"battleground/renderer"
	"battleground/simulation"
	"battleground/state"
	"math/rand"
	"time"
)

var winTitle string = "Go-SDL2 Events"
var winWidth, winHeight int32 = 800, 600

func main() {
	renderer.Init()
	defer renderer.Destroy()

	world := state.NewWorld()
	objManager := simulation.NewObjectManager(world)
	collisionManager := simulation.NewCollisionManager(world)

	// world.NewPlayerAt(1, 600, 320)
	// world.NewPlayerAt(2, 300, 300)
	// world.NewPlayerAt(3, 40, 310)
	// world.NewPlayerAt(4, 100, 300)
	// world.NewPlayerAt(5, 700, 300)
	// world.NewPlayerAt(6, 300, 500)
	// world.NewPlayerAt(7, 300, 100)
	// world.Players[6].SetFacing(-math.Pi / 2)
	// world.Players[7].SetFacing(math.Pi / 2)
	// world.Players[7].SetVelocity(15)
	// world.Players[1].SetFacing(math.Pi)
	// world.Players[2].SetFacing(0)
	// world.Players[1].SetVelocity(9)
	// world.Players[3].SetVelocity(7)
	// world.Players[5].SetVelocity(2)
	// world.Players[2].SetVelocity(-1)
	// world.Players[5].SetVelocity(-1)
	id := uint32(0)

	for _ = range time.Tick(20 * time.Millisecond) {
		if id < 50 {
			id++
			world.NewPlayerAt(id, 300, 300)
			world.Players[id].SetFacing((rand.Float64() - 0.5) * 3)
			world.Players[id].SetVelocity(rand.Float64()*20 + 1)
		}

		objManager.UpdateObjects()
		collisionManager.ResolveCollisions()
		renderer.Render(world)
	}

}
