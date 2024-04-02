package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Name    string
	Message string
}

type hub struct {
	clients               map[*websocket.Conn]bool
	clientRegisterChannel chan *websocket.Conn
	clientRemovalChannel  chan *websocket.Conn
	broadcastMessage      chan Message
}

func (h *hub) run() {
	for {
		select {
		case conn := <-h.clientRegisterChannel:
			h.clients[conn] = true
		case conn := <-h.clientRemovalChannel:
			delete(h.clients, conn)
		case message := <-h.broadcastMessage:
			for conn := range h.clients {
				_ = conn.WriteJSON(message)
			}
		}
	}
}

func main() {
	h := &hub{
		clients:               make(map[*websocket.Conn]bool),
		clientRegisterChannel: make(chan *websocket.Conn),
		clientRemovalChannel:  make(chan *websocket.Conn),
		broadcastMessage:      make(chan Message),
	}

	go h.run()

	app := fiber.New()
	app.Use("/ws", AllowUpgrade)
	app.Use("/ws/chat", websocket.New(Chat(h)))

	_ = app.Listen(":9000")
}

func AllowUpgrade(ctx *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(ctx) {
		return ctx.Next()
	}

	return fiber.ErrUpgradeRequired
}

func Chat(h *hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {

		defer func() {
			h.clientRemovalChannel <- conn
			_ = conn.Close()
		}()

		name := conn.Query("name")
		if name == "" {
			return
		}
		h.clientRegisterChannel <- conn

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				return
			}

			if messageType == websocket.TextMessage {
				h.broadcastMessage <- Message{
					Name:    name,
					Message: string(message),
				}
			}
		}
	}
}
