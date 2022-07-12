package game

import (
	"battleground/events"
	"battleground/gameticker"
	"battleground/simulation"
	"battleground/state"
	"math/rand"
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

	world.NewPlayerAt(1, 400, 350)
	world.Players[1].SetFacing((rand.Float64()) * 2 * 3.1415)
	world.Players[1].SetVelocity(rand.Float64()*5 + 1)
	world.Players[1].SetAngularVelocity((rand.Float64() - 0.5) / 6)

	eventHub := events.NewEventHub()
	physicTicker := gameticker.NewPhysicsTicker(eventHub)
	updater := simulation.NewUpdater(world, eventHub)
	eventHub.RegisterTimeTickListener(updater)
	// TODO: run event loop here
	go physicTicker.Run()
	go eventHub.RunEventLoop()

	return &GameRoom{
		roomID:       1,
		world:        world,
		updater:      updater,
		eventHub:     eventHub,
		clients:      make(map[uint32]*Client),
		sender:       NewSender(),
		physicTicker: physicTicker,
	}
}

//for testing
func (room *GameRoom) World() *state.World {
	return room.world
}
