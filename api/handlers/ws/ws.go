package ws

import (
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type message struct {
	data []byte
	room string
}

// serveWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request) error {
	ws, err := upgrader.Upgrade(w, r, nil)
	vars := r.URL.Query()
	rooms, ok := vars["room"]
	log.Println(rooms)
	var room string
	if ok && len(rooms) >= 1 {
		room = rooms[0]
	}
	if err != nil {
		return err
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, room}
	h.register <- s
	go s.writePump()
	// XXX なぜもとのやつに go がない?
	go s.readPump()
	return nil
}
