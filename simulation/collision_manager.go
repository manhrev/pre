package simulation

import (
	"battleground/constants"
	"battleground/state"
	"battleground/types"
	"math"
)

type CollisionManager struct {
	world *state.World
}

func NewCollisionManager(world *state.World) *CollisionManager {
	return &CollisionManager{world: world}
}

func (manager *CollisionManager) ResolveCollisionsOnMap() {
	for _, object := range manager.world.Objects {
		x, y := manager.world.Map.TileContain(object)
		//fmt.Println(manager.world.Map.GetTilesAt(int(x), int(y)))
		if manager.world.Map.GetTilesAt(int(x), int(y)) != 1 {
			object.SetFacing(object.Facing() + 2.1415)
		}
	}
}

func (manager *CollisionManager) ResolveCollisions() {
	collided := make(map[state.Object]bool)
	for _, object := range manager.world.Objects {
		collided[object] = true
		for _, otherObject := range manager.world.Objects {
			if !collided[otherObject] && object.DetectCollision(otherObject) {

				// Pull 2 object apart
				vObjtoOther := &types.Vector{
					X: float64(otherObject.Position().X - object.Position().X),
					Y: float64(otherObject.Position().Y - object.Position().Y),
				}
				// object.SetFacing(object.Facing() + math.Pi)
				// otherObject.SetFacing(object.Facing() + math.Pi)

				vObjtoOther = vObjtoOther.Normalize()
				dist := object.DistanceTo(otherObject)
				object.SetPosition(object.Position().Add(vObjtoOther.Multiply(constants.PlayerSize - dist/2).Rotate(math.Pi)))
				otherObject.SetPosition(otherObject.Position().Add(vObjtoOther.Multiply(constants.PlayerSize - dist/2)))
			}
		}
	}
}
