package gameticker

import (
	"battleground/constants"
	"battleground/events"
	"time"
)

type PhysicsTicker struct {
	eventHub *events.EventHub
}

func NewPhysicsTicker(eventHub *events.EventHub) *PhysicsTicker {
	return &PhysicsTicker{
		eventHub: eventHub,
	}
}

func (ticker *PhysicsTicker) Run() {
	var i uint32
	i = 0
	// remember when to stop a physic ticker goroutine
	for range time.Tick(constants.PhysicFrameDuration) {
		//TODO: send time tick event
		ticker.eventHub.FireEvent(events.TimeTick{
			FrameId: i,
		})
		i++
	}
}
