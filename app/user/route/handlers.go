package route

import (
	"encoding/json"
	"net/http"
)

type handler struct{}

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	user := user{"Ramadani", "email.ramadani@gmail.com"}
	json, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// NewHandler user
func NewHandler() *handler {
	return &handler{}
}
