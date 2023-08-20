package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"

	// Stylish stuff
	"github.com/charmbracelet/log"

	// Websockets
	"github.com/gorilla/websocket"
)

type App struct {
	UserHandler *UserHandler
	clients     map[*websocket.Conn]bool
	broadcast   chan []byte
	upgrader    websocket.Upgrader
}

type Player struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameState struct {
	Players []Player `json:"players"`
}

func (s *App) handlePlayerMovement(conn *websocket.Conn) {
	var player Player
	player.X = 0
	player.Y = 0

	for {
		// Read the next message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Error("Failed to read message:", err)
			break
		}

		// Update the player's position based on the received message
		switch string(message) {
		case "up":
			player.Y--
		case "down":
			player.Y++
		case "left":
			player.X--
		case "right":
			player.X++
		}

		// Create a JSON representation of the game state
		gameState := GameState{
			Players: []Player{player},
		}
		gameStateData, err := json.Marshal(gameState)
		if err != nil {
			log.Error("Failed to marshal game state:", err)
			break
		}

		// Broadcast the game state to all clients
		s.broadcast <- gameStateData
	}
}

func (s *App) handleBroadcasts() {
	for {
		// Read the next message from the broadcast channel
		message := <-s.broadcast

		// Broadcast the message to all clients
		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Error("Failed to write message:", err)
			}
		}
	}
}

var upgrader = websocket.Upgrader{}

// Define the serve function
func (s *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	// Start the user handler should the requested user be found
	case "user":
		s.UserHandler.Handle(res, req)
	// Return a `Not Found` if the user is not found
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

// Run the server
func (s *App) Run() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the connection to a WebSocket connection
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error("Failed to upgrade connection:", err)
			return
		}

		// Add the client connection to the clients map
		s.clients[conn] = true

		// Log when a client connects
		log.Info("Client connected:", conn.RemoteAddr())

		// Allow the server to handle player movement
		go s.handlePlayerMovement(conn)

		// Close the connection and remove it from the clients map
		defer func() {
			// Log when a client disconnects
			log.Info("Client disconnected:", conn.RemoteAddr())

			conn.Close()
			delete(s.clients, conn)
		}()

		// Handle incoming messages
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Error("Failed to read message:", err)
				break
			}

			// Broadcast the received message to all clients
			s.broadcast <- message
		}
	})

	go s.handleBroadcasts()

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Fatal(err)
		}

		for {

			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
}

func main() {
	// Make the jwt_secret file within the server configuration directory
	makeSecret()

	// Initialise the user handler
	user_handler, err := NewUserHandler()

	// Log any errors
	if err != nil {
		log.Fatal(err)
	}

	a := &App{
		UserHandler: user_handler,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "7741"
	}
	// Log that the program has successfully started listening to the port
	log.Info(fmt.Sprintf("Ambition backend listening to port %v", port))
	http.ListenAndServe(":"+port, a)

}
