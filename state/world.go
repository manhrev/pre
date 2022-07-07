package state

import (
	"battleground/types"
)

type World struct {
	Objects        map[uint32]Object
	Players        map[uint32]*Player
	PhysicsFrameID uint32
	Map            *Map
}

func NewWorld() *World {
	return &World{
		Objects:        make(map[uint32]Object),
		Players:        make(map[uint32]*Player),
		PhysicsFrameID: 0,
		Map:            NewMap(),
	}
}

func (world *World) NewPlayer(clientID uint32) *Player {
	player := NewPlayer(clientID, &types.Point{X: 0, Y: 0})
	world.AddPlayer(clientID, player)
	return player
}

func (world *World) NewPlayerAt(clientID uint32, x int32, y int32) *Player {
	player := NewPlayer(clientID, &types.Point{X: x, Y: y})
	world.AddPlayer(clientID, player)
	return player
}

func (world *World) AddPlayer(clientID uint32, player *Player) {
	world.Players[clientID] = player
	world.Objects[clientID] = player
}

func (world *World) RemoveObject(id uint32) {
	delete(world.Objects, id)
	delete(world.Players, id)

}
