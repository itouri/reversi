package ws

import (
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

func parse(key string, vars map[string][]string) string {
	vals, ok := vars[key]
	var val string
	if ok && len(vals) >= 1 {
		val = vals[0]
	}
	return val
}

// serveWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request) error {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	vars := r.URL.Query()
	room_id := parse("room_id", vars)
	// player_name := parse("player_name", vars)

	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, room_id}
	h.register <- s
	go s.writePump()
	// XXX なぜもとのやつに go がない?
	go s.readPump()
	return nil
}
