package user

import (
	"encoding/json"
	"net/http"

	"github.com/furqanrupom/sqlx-macro/pkg"
	"github.com/furqanrupom/sqlx-macro/util"
)

type LoginData struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type ResponseData struct {
	AccessToken string `json:"access_token"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData LoginData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginData)
	if err != nil {
		util.SendError(w,http.StatusInternalServerError,"Invalid json data!")
		return
	}
	userData,err := h.svc.Login(loginData.Email,loginData.Password)
	if err != nil {
		util.SendError(w,http.StatusBadRequest,"User Login Failed",err)
		return
	}
	token,err := pkg.CreateToken(pkg.Claims{
		Sub:userData.ID,
		Name:userData.Name,
		Email:userData.Email,
	},h.config.AuthConfig.SecretToken)
	if err !=  nil {
		util.SendError(w,http.StatusBadRequest,"Token creation failed!")
		return
	}
	util.SendData(w,http.StatusOK,&ResponseData{AccessToken: token})
}