package communication

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	id uint32
	ws *websocket.Conn
	// ch chan

}
