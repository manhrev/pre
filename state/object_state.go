package state

import (
	"battleground/constants"
	"battleground/types"
	"math"
	"time"
)

type ObjectState struct {
	id              uint32
	position        *types.Point
	velocity        float64 //0 to 2pi
	facing          float64
	angularVelocity float64
	//hp       uint32

	spawnTimestamp time.Time
}

func NewObjectState(ID uint32, position *types.Point, velocity float64) *ObjectState {
	return &ObjectState{
		id:             ID,
		position:       position,
		velocity:       velocity,
		facing:         0.0,
		spawnTimestamp: time.Now(),
	}
}

func (objectState *ObjectState) Id() uint32 {
	return objectState.id
}

func (objectState *ObjectState) Position() *types.Point {
	return objectState.position
}

func (objectState *ObjectState) VelocityVect() *types.Vector {
	return &types.Vector{
		X: objectState.Velocity() * math.Cos(objectState.Facing()),
		Y: objectState.Velocity() * math.Sin(objectState.Facing()),
	}
}

func (objectState *ObjectState) SetPosition(position *types.Point) {
	objectState.position = position
}

func (objectState *ObjectState) Velocity() float64 {
	return objectState.velocity
}

func (objectState *ObjectState) SetVelocity(velocity float64) {
	objectState.velocity = velocity
}

func (objectState *ObjectState) AngularVelocity() float64 {
	return objectState.angularVelocity
}

func (objectState *ObjectState) SetAngularVelocity(angularVelocity float64) {
	objectState.angularVelocity = angularVelocity
}

func (objectState *ObjectState) Facing() float64 {
	return objectState.facing
}

func (objectState *ObjectState) SetFacing(facing float64) {
	objectState.facing = facing
}

func (object *ObjectState) DetectCollision(other Object) bool {
	v := types.Point{X: object.Position().X - other.Position().X, Y: object.Position().Y - other.Position().Y}
	dist := v.Length()
	return dist < 2*constants.PlayerSize
}

func (object *ObjectState) DistanceTo(other Object) float64 {
	v := types.Point{X: object.Position().X - other.Position().X, Y: object.Position().Y - other.Position().Y}
	dist := v.Length()
	return dist
}
