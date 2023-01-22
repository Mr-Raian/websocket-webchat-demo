package handlers

import (
	"context"
	"net/http"
	"time"
	"websocket-webchat-demo/internal/ent"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	clients   = make(map[*websocket.Conn]bool) // connected clients
	broadcast = make(chan *ent.Message)        // broadcast channel
	upgrader  = websocket.Upgrader{}
	DB        = &ent.Client{}
)

func SendMessage(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// Register client
	clients[ws] = true

	for {
		// Read message from client
		message := &ent.Message{}
		err := ws.ReadJSON(message)
		if err != nil {
			c.Logger().Error(err)
			delete(clients, ws)
			break
		}

		// Save in DB
		_, err = DB.Message.Create().
			SetUsername(message.Username).
			SetMessage(message.Message).
			SetSentAt(time.Now()).Save(context.Background())
		if err != nil {
			c.Logger().Warnf("Error saving message in the DB: %s", err)
		}

		// Send message to all clients
		broadcast <- message
	}

	return nil
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
