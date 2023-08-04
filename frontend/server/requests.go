package server

import (
	"os"
	"errors"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"
	
	// Needed for server url
	_ "github.com/joho/godotenv/autoload"
)

var server_address = os.Getenv("BACKEND_URL")

func GetUserJwtToken(name, password string) (string, error) {
	if server_address == "" {
		return "", errors.New("no server url provided")
	}
	
	post_body, _ := json.Marshal(map[string]string{
		"name": name,
		"password": password,
	})
	resp_body := bytes.NewBuffer(post_body)
	
	resp, err := http.Post(server_address + "/user", "application/json", resp_body)
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