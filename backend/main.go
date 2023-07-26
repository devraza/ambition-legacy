package main

import (
	"log"
	"net/http"
	"os"
)

type App struct {
	UserHandler *UserHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "user":
		h.UserHandler.Handle(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func main() {
	user_handler, err := NewUserHandler()
	if err != nil {
		log.Fatalln(err)
	}

	a := &App{
		UserHandler: user_handler,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "7741"
	}
	log.Println("Ambition going strong at port 7741")
	http.ListenAndServe(":"+port, a)
}
