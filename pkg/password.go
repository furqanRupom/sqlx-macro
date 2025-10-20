package pkg

import "github.com/furqanrupom/sqlx-macro/config"

func HashPassword(passwordConfig *config.PasswordConfig, password string) (string, error) {
	if len(password) == 0 {
		return "", nil
	}
	return "", nil
}
func ComparePassword() {}