package server

import (
	"battleground/game"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Server handle new ws connect, find a game room for a client (or create new room)
type Server struct {
	gameRooms map[uint32]*game.GameRoom
	upgrader  *websocket.Upgrader
}

func NewServer() *Server {
	server := &Server{
		gameRooms: make(map[uint32]*game.GameRoom),
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return false
			}},
	}

	server.addGameRoom(game.NewGameRoom())
	return server

}

func (s *Server) addGameRoom(room *game.GameRoom) {
	s.gameRooms[room.RoomID()] = room
}
func (s *Server) Listen() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		// TODO: check if client exist

		//find room for client
		room := s.gameRooms[1]

		//new client
		client := game.NewClient(conn, room)
		room.AddClient(client)
		client.Listen()

	})
}
