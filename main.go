package main

import (
	"battleground/game"
	"battleground/renderer"
	"battleground/simulation"
	"battleground/state"
	"math/rand"
	"time"
)

func maind() {
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
		if id%30 == 0 {

			world.NewPlayerAt(id, 400, 350)
			world.Players[id].SetFacing((rand.Float64()) * 2 * 3.1415)
			world.Players[id].SetVelocity(rand.Float64()*15 + 1)
			world.Players[id].SetAngularVelocity((rand.Float64() - 0.5) / 6)
			//world.Players[id].SetVelocity(0)
		}
		id++
		objManager.UpdateObjects()
		collisionManager.ResolveCollisionsOnMap()
		collisionManager.ResolveCollisions()
		collisionManager.ResolveCollisionsOnMap()
		//println(len(world.Objects))
		if !renderer.Render(world) {
			return
		}

	}

}

func main() {
	renderer.Init()
	defer renderer.Destroy()

	room := game.NewGameRoom()

	world := room.World()

	// objManager := simulation.NewObjectManager(world)
	// collisionManager := simulation.NewCollisionManager(world)

	// id := uint32(0)

	for _ = range time.Tick(20 * time.Millisecond) {
		// if id%30 == 0 {

		// 	world.NewPlayerAt(id, 400, 350)
		// 	world.Players[id].SetFacing((rand.Float64()) * 2 * 3.1415)
		// 	world.Players[id].SetVelocity(rand.Float64()*15 + 1)
		// 	world.Players[id].SetAngularVelocity((rand.Float64() - 0.5) / 6)
		// 	//world.Players[id].SetVelocity(0)
		// }
		// id++
		// objManager.UpdateObjects()
		// collisionManager.ResolveCollisionsOnMap()
		// collisionManager.ResolveCollisions()
		// collisionManager.ResolveCollisionsOnMap()
		// //println(len(world.Objects))
		if !renderer.Render(world) {
			return
		}

	}

}
