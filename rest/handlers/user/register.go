package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/furqanrupom/sqlx-macro/domain"
	"github.com/furqanrupom/sqlx-macro/util"
)

type User struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request){
	var reqData User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		util.SendError(w,http.StatusInternalServerError,"Please give me valid json!")
	}
	createUser,err := h.svc.Register(domain.User{
		Name:reqData.Name,
		Email:reqData.Email,
		Password: reqData.Password,
	})
	if err != nil {
		log.Println(err)
		util.SendError(w,http.StatusBadRequest,"User creation failed",err)
		return
	}
	util.SendData(w,http.StatusCreated,createUser)
}