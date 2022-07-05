package state

import (
	"battleground/types"
)

type Player struct {
	ObjectState
	Boost bool
}

func NewPlayer(clientId uint32, initialPosition *types.Point) *Player {
	objectState := NewObjectState(clientId, initialPosition, 0)

	//for testing
	objectState.SetVelocity(5)

	return &Player{
		ObjectState: *objectState,
	}
}
