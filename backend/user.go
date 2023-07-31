package main

import (
	// The standard stuff
	"errors"
	"fmt"
	"io"
	"os"

	// Encryption
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	// SQL databasing
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Define the user request struct
type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (ur *UserRequest) Parse(req *http.Request) error {
	// Can't unmarshal the actual req.Body so must read first
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &ur); err != nil {
		return err
	}
	return nil
}

// Define the user handler struct
type UserHandler struct {
	db         *sql.DB
	jwt_secret []byte
}

// Define the function to create user handlers
func NewUserHandler() (*UserHandler, error) {
	// Initialise the database using the database file
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}

	// Define the JSON web token
	jwt_secret_str := os.Getenv("JWT_SECRET")
	// Return any errors
	if jwt_secret_str == "" {
		return nil, errors.New("no JWT_SECRET provided in .env")
	}
	jwt_secret := []byte(jwt_secret_str)

	// Return the user handler struct
	return &UserHandler{
		db:         db,
		jwt_secret: jwt_secret,
	}, nil
}

// JWT Utilities
func (h *UserHandler) ParseUserToken(token_string string) (*jwt.Token, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("jwt: incorrect token signing method")
		}
		return h.jwt_secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (h *UserHandler) GenerateUserToken(name, pwdhash string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["pwdhash"] = pwdhash

	token_string, err := token.SignedString(h.jwt_secret)
	if err != nil {
		return "", err
	}
	return token_string, nil
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
	user_request := new(UserRequest)
	if err := user_request.Parse(req); err != nil {
		http.Error(res, fmt.Sprintf("user: failed to parse request (%s)", err), http.StatusBadRequest)
		return
	}
	name := user_request.Name
	password := []byte(user_request.Password)

	// Password checks
	row := h.db.QueryRow("SELECT pwdhash FROM users WHERE name=?", name)
	var db_pwdhash string

	if err := row.Scan(&db_pwdhash); err != nil {
		// If no user is found with the requested name, create the user
		if errors.Is(err, sql.ErrNoRows) {
			pwdhash_bytes, err := bcrypt.GenerateFromPassword(password, 12)
			// Log any errors
			if err != nil {
				http.Error(res, fmt.Sprintf("user: failed to generate password hash (%s)", err), http.StatusInternalServerError)
				return
			}
			db_pwdhash = string(pwdhash_bytes)

			_, err = h.db.Exec("INSERT INTO users VALUES (?,?)", name, db_pwdhash)
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

	// JWT Generation
	token_string, err := h.GenerateUserToken(name, db_pwdhash)
	if err != nil {
		http.Error(res, fmt.Sprintf("jwt: failed to generate token (%s)", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(res, "%s", token_string)
}

func (h *UserHandler) updateUser(res http.ResponseWriter, req *http.Request) {
	user_request := new(UserRequest)
	if err := user_request.Parse(req); err != nil {
		http.Error(res, fmt.Sprintf("user: failed to parse request (%s)", err), http.StatusBadRequest)
		return
	}
	req_name := user_request.Name
	req_password := []byte(user_request.Password)

	if req.Header["Authorization"] == nil {
		http.Error(res, "jwt: missing token", http.StatusUnauthorized)
		return
	}

	token, err := h.ParseUserToken(req.Header["Token"][0])
	if err != nil {
		http.Error(res, fmt.Sprintf("jwt: error during parsing (%s)", err), http.StatusInternalServerError)
		return
	}
	if !token.Valid {
		http.Error(res, "jwt: invalid token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(res, "jwt: failed to get claims", http.StatusInternalServerError)
		return
	}

	claim_name := claims["name"].(string)
	claim_pwdhash := claims["pwdhash"].(string)

	var db_name, db_pwdhash string
	row := h.db.QueryRow("SELECT * FROM users WHERE name=?", claim_name)
	if err := row.Scan(&db_name, &db_pwdhash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(res, "user: authorized user doesn't exist", http.StatusBadRequest)
			return
		} else {
			http.Error(res, fmt.Sprintf("db: failed to find user (%s)", err), http.StatusInternalServerError)
			return
		}
	}

	if claim_pwdhash != db_pwdhash {
		http.Error(res, "user: invalid password", http.StatusForbidden)
		return
	}

	if req_name == claim_name && bcrypt.CompareHashAndPassword([]byte(claim_pwdhash), req_password) == nil {
		http.Error(res, "user: requested credentials are the same as current credentials", http.StatusBadRequest)
		return
	}

	req_pwdhash_bytes, err := bcrypt.GenerateFromPassword(req_password, 12)
	if err != nil {
		http.Error(res, fmt.Sprintf("user: failed to generate password hash (%s)", err), http.StatusInternalServerError)
		return
	}
	req_pwdhash := string(req_pwdhash_bytes)

	_, err = h.db.Exec("UPDATE users SET name=?, pwdhash=? WHERE name=? AND pwdhash=?",
		req_name, req_pwdhash, claim_name, claim_pwdhash)
	if err != nil {
		http.Error(res, fmt.Sprintf("db: failed to update user (%s)", err), http.StatusInternalServerError)
	}

	token_string, err := h.GenerateUserToken(req_name, req_pwdhash)
	if err != nil {
		http.Error(res, fmt.Sprintf("jwt: failed to generate token (%s)", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(res, "%s", token_string)
}

// TODO(midnadimple): Implement:
func (h *UserHandler) deleteUser(res http.ResponseWriter, req *http.Request) {
	token, err := h.ParseUserToken(req.Header["Token"][0])
	if err != nil {
		http.Error(res, fmt.Sprintf("jwt: error during parsing (%s)", err), http.StatusInternalServerError)
		return
	}
	if !token.Valid {
		http.Error(res, "jwt: invalid token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(res, "jwt: failed to get claims", http.StatusInternalServerError)
		return
	}

	claim_name := claims["name"].(string)
	claim_pwdhash := claims["pwdhash"].(string)

	var db_name, db_pwdhash string
	row := h.db.QueryRow("SELECT * FROM users WHERE name=?", claim_name)
	if err := row.Scan(&db_name, &db_pwdhash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(res, "user: authorized user doesn't exist", http.StatusBadRequest)
			return
		} else {
			http.Error(res, fmt.Sprintf("db: failed to find user (%s)", err), http.StatusInternalServerError)
			return
		}
	}

	if claim_pwdhash != db_pwdhash {
		http.Error(res, "user: invalid password", http.StatusForbidden)
		return
	}

	if _, err := h.db.Exec("DELETE FROM users WHERE name=? AND pwdhash=?", db_name, db_pwdhash); err != nil {
		http.Error(res, fmt.Sprintf("db: failed to delete user (%s)", err), http.StatusInternalServerError)
		return
	}

}
