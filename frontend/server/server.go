package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	// Stylish stuff
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

	// Needed for server url
	_ "github.com/joho/godotenv/autoload"
)

// Define a bold green style
var bold = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("2"))

var server_address = os.Getenv("BACKEND_URL")

func GetUserJwtToken(name, password string) (string, error) {
	if server_address == "" {
		log.Info(fmt.Sprintf("Make sure you've set the %s environment variable!", bold.Render("BACKEND_URL")))
		log.Fatal("No server URL provided!")
	}

	post_body, _ := json.Marshal(map[string]string{
		"name":     name,
		"password": password,
	})
	resp_body := bytes.NewBuffer(post_body)

	resp, err := http.Post(server_address+"/user", "application/json", resp_body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
