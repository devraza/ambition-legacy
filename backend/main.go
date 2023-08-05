package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"

	// Stylish stuff
	"github.com/charmbracelet/log"
)

type App struct {
	UserHandler *UserHandler
}

// Define the serve function
func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	// Start the user handler should the requested user be found
	case "user":
		h.UserHandler.Handle(res, req)
	// Return a `Not Found` if the user is not found
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

// Run the server
func main() {
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
