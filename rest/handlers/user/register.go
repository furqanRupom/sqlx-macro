package user

import "net/http"

type User struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request){}