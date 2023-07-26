package main

import (
	// The standard stuff
	"errors"
	"fmt"
	"io"

	// Encryption
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	// SQL databasing
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Define the user handler struct
type UserHandler struct {
	db         *sql.DB
	jwt_secret *ecdsa.PrivateKey
}

// Define the user request struct
type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Define the function to create user handlers
func NewUserHandler() (*UserHandler, error) {
	// Initialise the database using the database file
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}

	// Define the JSON web token
	jwt_secret, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// Return any errors
	if err != nil {
		return nil, err
	}

	// Return the user handler struct
	return &UserHandler{
		db:         db,
		jwt_secret: jwt_secret,
	}, nil
}

func (h *UserHandler) Handle(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		h.createUser(res, req)
	case "PUT":
		h.updateUser(res, req)
	case "DELETE":
		h.deleteUser(res, req)
	// Return an error message should an invalid method be used
	default:
		http.Error(res, "Only POST, PUT, and DELETE are valid methods", http.StatusMethodNotAllowed)
	}
}

// NOTE(midnadimple): This function could be considered to do too much stuff, but
// I think this is the best implementation
func (h *UserHandler) createUser(res http.ResponseWriter, req *http.Request) {
	// Can't unmarshal the actual req.Body so must read first
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, fmt.Sprintf("user: failed to read request (%s)", err), http.StatusBadRequest)
		return
	}

	var user_request UserRequest
	if err := json.Unmarshal(body, &user_request); err != nil {
		http.Error(res, "user: json request body doesn't match schema", http.StatusBadRequest)
		return
	}

	name := user_request.Name
	password := []byte(user_request.Password)

	// Password checks
	row := h.db.QueryRow("SELECT pwdhash FROM users WHERE name=?", name)
	var db_pwdhash string

	if err = row.Scan(&db_pwdhash); err != nil {
		// If no user is found with the requested name, create the user
		if errors.Is(err, sql.ErrNoRows) {
			pwdhash_bytes, err := bcrypt.GenerateFromPassword(password, 12)
			// Log any errors
			if err != nil {
				http.Error(res, fmt.Sprintf("user: failed to generate password hash (%s)", err), http.StatusInternalServerError)
				return
			}
			pwdhash := string(pwdhash_bytes)

			_, err = h.db.Exec("INSERT INTO users VALUES (?,?)", name, pwdhash)
			if err != nil {
				http.Error(res, fmt.Sprintf("db: failed to create user (%s)", err), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(res, fmt.Sprintf("db: failed to query row (%s)", err), http.StatusInternalServerError)
			return
		}
	} else if bcrypt.CompareHashAndPassword([]byte(db_pwdhash), password) != nil {
		http.Error(res, "User exists, but invalid password", http.StatusForbidden)
		return
	}

	// JWT generation
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["pwdhash"] = db_pwdhash

	token_string, err := token.SignedString(h.jwt_secret)
	if err != nil {
		http.Error(res, fmt.Sprintf("jwt: failed to generate token (%s)", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(res, "%s", token_string)
}

// TODO(midnadimple): implement:
func (h *UserHandler) updateUser(res http.ResponseWriter, req *http.Request) {}
func (h *UserHandler) deleteUser(res http.ResponseWriter, req *http.Request) {}
