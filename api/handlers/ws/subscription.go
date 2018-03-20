package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type subscription struct {
	conn *connection
	room string
}

// TODO ここに具体的な処理を書かない
func (s *subscription) rematch(msg []byte) {
	var data jsonReversi
	if err := json.Unmarshal(msg, &data); err != nil {
		log.Print(err)
	}
	// 中身が rematch なら stoneColor を各クライアントに送信
	if data.FuncName == "rematch" {
		i := 0
		colors := []string{"-1", "1"}
		connections := h.rooms[s.room]
		for c := range connections {
			sendMsg := &jsonReversi{
				FuncName: "stoneColor",
				Body:     colors[i],
			}
			sendJSON, err := json.Marshal(sendMsg)
			if err != nil {
				log.Print(err)
			}
			m := uniMessage{s.room, sendJSON, c}
			h.unicast <- m
			i++
		}
	}
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		h.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		// TODO ここに具体的な処理を書かない
		s.rematch(msg)

		m := message{s.room, msg}
		h.broadcast <- m
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	log.Println("writePump") //!!!
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
