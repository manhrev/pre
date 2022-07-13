package game

import (
	"battleground/events"
	"battleground/gameticker"
	"battleground/simulation"
	"battleground/state"
	"math"
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

	eventHub := events.NewEventHub()
	physicTicker := gameticker.NewPhysicsTicker(eventHub)
	updater := simulation.NewUpdater(world, eventHub)
	eventHub.RegisterTimeTickListener(updater)

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

// Start game room
func (room *GameRoom) Start() {
	go room.physicTicker.Run()
	go room.eventHub.RunEventLoop()

	// TODO: from room.clients -> create players in world

	// for testing
	room.world.NewPlayerAt(1, 400, 400)
	room.world.Players[1].SetFacing(math.Pi / 4)
	room.world.Players[1].SetVelocity(rand.Float64()/4 + 1)
	//room.world.Players[1].SetAngularVelocity((rand.Float64() - 0.5) / 6)
}

func (room *GameRoom) Stop() {
	// Kill all goroutine
}

//for testing
func (room *GameRoom) World() *state.World {
	return room.world
}
