package state

import "battleground/types"

type Object interface {
	Id() uint32

	Position() *types.Point
	VelocityVect() *types.Vector
	Velocity() float64
	Facing() float64

	//Hp() uint32

	SetPosition(*types.Point)
	SetVelocity(float64)
	SetFacing(float64)

	DetectCollision(other Object) bool
	DistanceTo(other Object) float64
}
