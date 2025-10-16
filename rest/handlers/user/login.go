package user

import "net/http"

type LoginData struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

}