package ws

import (
	"log"
	"net/http"
)

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	rooms map[string]map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan message

	unicast chan uniMessage

	// Register requests from the connections.
	register chan subscription

	// Unregister requests from connections.
	unregister chan subscription
}

var h = hub{
	broadcast:  make(chan message),
	unicast:    make(chan uniMessage),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]map[*connection]bool),
}

func RunHab() {
	go h.run()
}

func sendExit(roomID string, playerID string, playerName string) {
	url := "http://localhost:12345/api/v1/rooms/" + roomID + "/" + playerID + "/" + playerName
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println(err)
	}
	client := new(http.Client)
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}
}

func (h *hub) run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				// TODO 接続解除の処理はここにあるべきなのか?
				sendExit(s.room, s.player.ID, s.player.Name)
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		case m := <-h.unicast:
			select {
			case m.conn.send <- m.data:
				// TODO ここいる?
			}
		}
	}
}
