// for manage and update player
package simulation

import (
	"battleground/state"
)

type PlayerManager struct {
	world *state.World
}

func NewPlayerManager(world *state.World) *PlayerManager {
	return &PlayerManager{
		world: world,
	}
}

func (manager *PlayerManager) UpdatePlayer() {
	// for _, player := range manager.world.Players {

	// }
}
