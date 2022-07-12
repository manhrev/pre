package game

import (
	"battleground/events"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	channelBufferSize = 100
)

type Client struct {
	//id in a game
	id uint32

	username string
	ws       *websocket.Conn
	ch       chan *[]byte
	gameRoom *GameRoom

	// ch chan
	// send chan *pb.ServerMessage

}

func NewClient(ws *websocket.Conn, gameRoom *GameRoom) *Client {
	ch := make(chan *[]byte, channelBufferSize)

	return &Client{
		id:       1,
		username: "manage",
		ws:       ws,
		ch:       ch,
		gameRoom: gameRoom,
	}
}

func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Client) listenWrite() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := c.ws.Close()
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}()
	for {
		select {
		case message := <-c.ch:
			c.ws.SetWriteDeadline((time.Now().Add(writeWait)))

			w, err := c.ws.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(*message)

		case <-ticker.C:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) listenRead() {
	defer func() {
		// TODO: send userleft event -> remove client from map
		err := c.ws.Close()
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}()

	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.ws.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Println(message)

		// TODO: handle message, for example
		c.gameRoom.eventHub.FireEvent(events.TimeTick{})

	}
}
