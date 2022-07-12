package game

import (
	"battleground/events"
	"battleground/gameticker"
	"battleground/simulation"
	"battleground/state"
)

type GameRoom struct {
	// Running or waiting
	running bool

	roomID uint32

	// event hub
	eventHub *events.EventHub

	// Game state
	world *state.World

	// Simulation (updater)
	updater *simulation.Updater

	// Sender
	sender *Sender

	// PhysicsTicker
	physicTicker *gameticker.PhysicsTicker

	clients map[uint32]*Client
}

func (room *GameRoom) RoomID() uint32 {
	return room.roomID
}

func (room *GameRoom) AddClient(client *Client) {
	// Assign id for client
	client.id = 1

	// Client room
	client.gameRoom = room

	// Add client to room
	room.clients[client.id] = client
}

func NewGameRoom() *GameRoom {
	world := state.NewWorld()
	eventHub := events.NewEventHub()
	physicTicker := gameticker.NewPhysicsTicker(eventHub)

	// TODO: run event loop here
	go physicTicker.Run()

	return &GameRoom{
		roomID:       1,
		world:        world,
		updater:      simulation.NewUpdater(world, eventHub),
		eventHub:     eventHub,
		clients:      make(map[uint32]*Client),
		sender:       NewSender(),
		physicTicker: physicTicker,
	}
}
