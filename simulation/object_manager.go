// for simulating movement of all object (include player)
package simulation

import (
	"battleground/state"
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

		// position update
		object.SetPosition(object.Position().Add(object.VelocityVect()))

		// apply angularVelocity
		object.SetFacing(object.Facing() + object.AngularVelocity())

	}
}
