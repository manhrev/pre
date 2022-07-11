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
		switch manager.world.Map.GetTilesAt(int(x), int(y)) {
		case 0:
			manager.world.RemoveObject(object.Id())
			return

		case 1:

			vX := &types.Vector{
				X: -1,
				Y: 0,
			}
			vY := &types.Vector{
				X: 0,
				Y: -1,
			}
			// Check collision with wall
			if manager.world.Map.GetTilesAt(int(x+1), int(y)) == 2 {
				vObjToTile := &types.Vector{
					X: float64(x+1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().X,
					Y: float64(y)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().Y,
				}
				if vObjToTile.Length() < math.Sqrt(2)*constants.MapTileSize/2+constants.PlayerSize {
					distObjToTileX := math.Abs(vObjToTile.Length() * math.Cos(vObjToTile.Radians()))
					distObjToTileY := math.Abs(vObjToTile.Length() * math.Sin(vObjToTile.Radians()))
					// Collision occured
					if distObjToTileX-constants.PlayerSize < constants.MapTileSize/2 && distObjToTileY-constants.PlayerSize < constants.MapTileSize/2 {

						if distObjToTileX > distObjToTileY {
							object.SetPosition(object.Position().Add(vX.Multiply(-distObjToTileX + constants.MapTileSize/2 + constants.PlayerSize)))
						} else {
							object.SetPosition(object.Position().Add(vY.Multiply(-distObjToTileY + constants.MapTileSize/2 + constants.PlayerSize)))
						}
					}
				}

			}
			if manager.world.Map.GetTilesAt(int(x-1), int(y)) == 2 {
				//fmt.Println("left")
				vObjToTile := &types.Vector{
					X: float64(x-1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().X,
					Y: float64(y)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().Y,
				}
				if vObjToTile.Length() < math.Sqrt(2)*constants.MapTileSize/2+constants.PlayerSize {
					distObjToTileX := math.Abs(vObjToTile.Length() * math.Cos(vObjToTile.Radians()))
					distObjToTileY := math.Abs(vObjToTile.Length() * math.Sin(vObjToTile.Radians()))

					//println(distObjToTileX, distObjToTileY)

					// Collision occured
					if distObjToTileX-constants.PlayerSize < constants.MapTileSize/2 && distObjToTileY-constants.PlayerSize < constants.MapTileSize/2 {
						//fmt.Println("collision with left")
						if distObjToTileX > distObjToTileY {
							//fmt.Println("x far")
							object.SetPosition(object.Position().Add(vX.Multiply(-(-distObjToTileX + constants.MapTileSize/2 + constants.PlayerSize))))
						} else {
							//fmt.Println("y far")
							object.SetPosition(object.Position().Add(vY.Multiply(-distObjToTileY + constants.MapTileSize/2 + constants.PlayerSize)))
						}
					}
				}

			}
			if manager.world.Map.GetTilesAt(int(x), int(y-1)) == 2 {
				//println("top")
				vObjToTile := &types.Vector{
					X: float64(x)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().X,
					Y: float64(y-1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().Y,
				}
				if vObjToTile.Length() < math.Sqrt(2)*constants.MapTileSize/2+constants.PlayerSize {
					distObjToTileX := math.Abs(vObjToTile.Length() * math.Cos(vObjToTile.Radians()))
					distObjToTileY := math.Abs(vObjToTile.Length() * math.Sin(vObjToTile.Radians()))

					// Collision occured
					if distObjToTileX-constants.PlayerSize < constants.MapTileSize/2 && distObjToTileY-constants.PlayerSize < constants.MapTileSize/2 {
						if distObjToTileX > distObjToTileY {
							object.SetPosition(object.Position().Add(vX.Multiply(-distObjToTileX + constants.MapTileSize/2 + constants.PlayerSize)))
						} else {
							object.SetPosition(object.Position().Add(vY.Multiply(-(-distObjToTileY + constants.MapTileSize/2 + constants.PlayerSize))))
						}
					}
				}

			}
			if manager.world.Map.GetTilesAt(int(x), int(y+1)) == 2 {
				vObjToTile := &types.Vector{
					X: float64(x)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().X,
					Y: float64(y+1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().Y,
				}
				if vObjToTile.Length() < math.Sqrt(2)*constants.MapTileSize/2+constants.PlayerSize {
					distObjToTileX := math.Abs(vObjToTile.Length() * math.Cos(vObjToTile.Radians()))
					distObjToTileY := math.Abs(vObjToTile.Length() * math.Sin(vObjToTile.Radians()))

					// Collision occured
					if distObjToTileX-constants.PlayerSize < constants.MapTileSize/2 && distObjToTileY-constants.PlayerSize < constants.MapTileSize/2 {
						if distObjToTileX > distObjToTileY {
							object.SetPosition(object.Position().Add(vX.Multiply(-distObjToTileX + constants.MapTileSize/2 + constants.PlayerSize)))
						} else {
							object.SetPosition(object.Position().Add(vY.Multiply(-distObjToTileY + constants.MapTileSize/2 + constants.PlayerSize)))
						}
					}
				}

			}
			if manager.world.Map.GetTilesAt(int(x-1), int(y-1)) == 2 {
				vObjToTile := &types.Vector{
					X: float64(x-1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().X,
					Y: float64(y-1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().Y,
				}
				if vObjToTile.Length() < math.Sqrt(2)*constants.MapTileSize/2+constants.PlayerSize {
					distObjToTileX := math.Abs(vObjToTile.Length() * math.Cos(vObjToTile.Radians()))
					distObjToTileY := math.Abs(vObjToTile.Length() * math.Sin(vObjToTile.Radians()))

					// Collision occured
					if distObjToTileX-constants.PlayerSize < constants.MapTileSize/2 && distObjToTileY-constants.PlayerSize < constants.MapTileSize/2 {
						if distObjToTileX > distObjToTileY {
							object.SetPosition(object.Position().Add(vX.Multiply(-(-distObjToTileX + constants.MapTileSize/2 + constants.PlayerSize))))
						} else {
							object.SetPosition(object.Position().Add(vY.Multiply(-(-distObjToTileY + constants.MapTileSize/2 + constants.PlayerSize))))
						}
					}
				}

			}
			if manager.world.Map.GetTilesAt(int(x-1), int(y+1)) == 2 {
				vObjToTile := &types.Vector{
					X: float64(x-1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().X,
					Y: float64(y+1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().Y,
				}
				if vObjToTile.Length() < math.Sqrt(2)*constants.MapTileSize/2+constants.PlayerSize {
					distObjToTileX := math.Abs(vObjToTile.Length() * math.Cos(vObjToTile.Radians()))
					distObjToTileY := math.Abs(vObjToTile.Length() * math.Sin(vObjToTile.Radians()))

					// Collision occured
					if distObjToTileX-constants.PlayerSize < constants.MapTileSize/2 && distObjToTileY-constants.PlayerSize < constants.MapTileSize/2 {
						if distObjToTileX > distObjToTileY {
							object.SetPosition(object.Position().Add(vX.Multiply(-(-distObjToTileX + constants.MapTileSize/2 + constants.PlayerSize))))
						} else {
							object.SetPosition(object.Position().Add(vY.Multiply(-distObjToTileY + constants.MapTileSize/2 + constants.PlayerSize)))
						}
					}
				}

			}
			if manager.world.Map.GetTilesAt(int(x+1), int(y-1)) == 2 {
				vObjToTile := &types.Vector{
					X: float64(x+1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().X,
					Y: float64(y-1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().Y,
				}
				if vObjToTile.Length() < math.Sqrt(2)*constants.MapTileSize/2+constants.PlayerSize {
					distObjToTileX := math.Abs(vObjToTile.Length() * math.Cos(vObjToTile.Radians()))
					distObjToTileY := math.Abs(vObjToTile.Length() * math.Sin(vObjToTile.Radians()))

					// Collision occured
					if distObjToTileX-constants.PlayerSize < constants.MapTileSize/2 && distObjToTileY-constants.PlayerSize < constants.MapTileSize/2 {
						if distObjToTileX > distObjToTileY {
							object.SetPosition(object.Position().Add(vX.Multiply(-distObjToTileX + constants.MapTileSize/2 + constants.PlayerSize)))
						} else {
							object.SetPosition(object.Position().Add(vY.Multiply(-(-distObjToTileY + constants.MapTileSize/2 + constants.PlayerSize))))
						}
					}
				}

			}
			if manager.world.Map.GetTilesAt(int(x+1), int(y+1)) == 2 {
				vObjToTile := &types.Vector{
					X: float64(x+1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().X,
					Y: float64(y+1)*constants.MapTileSize + constants.MapTileSize/2 - object.Position().Y,
				}
				if vObjToTile.Length() < math.Sqrt(2)*constants.MapTileSize/2+constants.PlayerSize {
					distObjToTileX := math.Abs(vObjToTile.Length() * math.Cos(vObjToTile.Radians()))
					distObjToTileY := math.Abs(vObjToTile.Length() * math.Sin(vObjToTile.Radians()))

					// Collision occured

					if distObjToTileX-constants.PlayerSize < constants.MapTileSize/2 && distObjToTileY-constants.PlayerSize < constants.MapTileSize/2 {
						if distObjToTileX > distObjToTileY {
							object.SetPosition(object.Position().Add(vX.Multiply(-distObjToTileX + constants.MapTileSize/2 + constants.PlayerSize)))
						} else {
							object.SetPosition(object.Position().Add(vY.Multiply(-distObjToTileY + constants.MapTileSize/2 + constants.PlayerSize)))
						}
					}
				}

			}
			break

		case 2:
			// May be push player to nearest platform
			//object.SetFacing(object.Facing() + 2.1415)
			break
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
