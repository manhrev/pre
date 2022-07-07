package simulation

import "battleground/state"

type Updater struct {
	world            *state.World
	playerManager    *PlayerManager
	objectManager    *ObjectManager
	collisionManager *CollisionManager
}
