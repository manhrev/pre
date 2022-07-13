package game

import "battleground/state"

type Sender struct {
	room  *GameRoom
	world *state.World
}

func NewSender(room *GameRoom, world *state.World) *Sender {
	return &Sender{
		room:  room,
		world: world,
	}
}
