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
	for _, object := range manage.world.Objects {
		// for testing
		if object.DistanceTo(state.NewObjectState(100, types.NewPoint(400, 300), 0)) >= 350 {

			object.SetFacing(object.Facing() + math.Pi)
			object.SetPosition(object.Position().Add(object.VelocityVect().Multiply(4)))
		}
		// position update
		object.SetPosition(object.Position().Add(object.VelocityVect()))

		// for testing
		// object.SetFacing(
		// 	object.Facing() + (rand.Float64() - 0.5),
		// )
	}
}
