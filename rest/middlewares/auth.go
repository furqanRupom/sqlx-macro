package middlewares

import (
	"net/http"
	"strings"

	"github.com/furqanrupom/sqlx-macro/util"
	"github.com/furqanrupom/sqlx-macro/pkg"
)

func (m *Middlewares) AuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			util.SendError(w, http.StatusForbidden, "Unauthorized!")
			return
		}
		headerSplit := strings.Split(header, " ")
		if len(headerSplit) != 2 {
			util.SendError(w, http.StatusBadRequest, "Invalid token!")
			return
		}
		tokenSign := headerSplit[1]
		_, err := pkg.VerifyToken(tokenSign, m.config.AuthConfig.SecretToken)
		if err != nil {
			util.SendError(w, http.StatusBadRequest, "Token verify failed!")
			return
		}
		next.ServeHTTP(w, r)
	})
}
