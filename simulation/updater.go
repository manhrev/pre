package simulation

import (
	"battleground/events"
	"battleground/state"
	"math/rand"
)

type Updater struct {
	eventHub         *events.EventHub
	world            *state.World
	playerManager    *PlayerManager
	objectManager    *ObjectManager
	collisionManager *CollisionManager
}

func NewUpdater(world *state.World, eventHub *events.EventHub) *Updater {
	return &Updater{
		eventHub:         eventHub,
		world:            world,
		playerManager:    NewPlayerManager(world),
		objectManager:    NewObjectManager(world),
		collisionManager: NewCollisionManager(world),
	}
}

func (updater *Updater) updatePhysics() {
	updater.objectManager.UpdateObjects()
	updater.collisionManager.ResolveCollisionsOnMap()
	updater.collisionManager.ResolveCollisions()
	updater.collisionManager.ResolveCollisionsOnMap()
}

func (updater *Updater) HandleTimeTick(*events.TimeTick) {
	updater.updatePhysics()
}

func (updater *Updater) HandleUserInput(*events.UserInput) {
	// TODO
}

func (updater *Updater) HandleUserJoined(event *events.UserJoined) {
	// TODO set position for user
	id := event.ClientID
	updater.world.NewPlayerAt(id, 400, 400)
	updater.world.Players[id].SetFacing((rand.Float64()) * 2 * 3.1415)
	updater.world.Players[id].SetVelocity(rand.Float64()*5 + 1)
	updater.world.Players[id].SetAngularVelocity((rand.Float64() - 0.5) / 6)

}

// HandleUserLeft
// HandleUserUserJoined
