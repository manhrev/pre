// for simulating movement of all object (include player)
package simulation

import (
	"battleground/state"
	"battleground/types"
	"math"
)

type ObjectManager struct {
	world *state.World
}

func NewObjectManager(world *state.World) *ObjectManager {
	return &ObjectManager{
		world: world,
	}
}

func (manage *ObjectManager) UpdateObjects() {
	// Update object position
	for _, object := range manage.world.Objects {
		// for testing
		if object.DistanceTo(state.NewObjectState(100, types.NewPoint(450, 450), 0)) >= 900 {

			object.SetFacing(object.Facing() + math.Pi)
			object.SetPosition(object.Position().Add(object.VelocityVect().Multiply(2)))
		}
		// position update
		object.SetPosition(object.Position().Add(object.VelocityVect()))

	}
}
