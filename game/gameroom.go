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
	updater := simulation.NewUpdater(world, eventHub)
	sender := NewSender(nil, world) // room pointer in sender is nil, update this pointer before this func returns
	eventHub.RegisterTimeTickListener(updater)
	eventHub.RegisterUserJoinedListener(updater)

	room := &GameRoom{
		running:      false,
		roomID:       1, //room id ?
		world:        world,
		updater:      updater,
		eventHub:     eventHub,
		clients:      make(map[uint32]*Client),
		sender:       sender,
		physicTicker: physicTicker,
	}
	sender.room = room
	return room
}

// Start game room
func (room *GameRoom) Start() {

	// For testing
	room.clients[1] = &Client{username: "dfs"}
	room.clients[2] = &Client{username: "lol"}

	// From room.clients -> create players in world
	for clientId, client := range room.clients {
		room.eventHub.FireEvent(&events.UserJoined{
			ClientID: clientId,
			UserName: client.username,
		})
	}

	// Run go routines
	go room.physicTicker.Run()
	go room.eventHub.RunEventLoop()

}

func (room *GameRoom) Stop() {
	// Kill all goroutine
}

//for rendering
func (room *GameRoom) World() *state.World {
	return room.world
}
