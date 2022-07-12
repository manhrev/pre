package simulation

import (
	"battleground/events"
	"battleground/state"
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
