package security

import (
	"blog-chi-gorm/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func Authorization(r *http.Request) (string, error) {
	keys := r.URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token, errors.New("token tidak ditemukan")
	}

	bearerToken := r.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		tokenString := strings.Split(bearerToken, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Config.SECRETKEY), nil
		})

		if err != nil {
			return "", err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			b, err := json.MarshalIndent(claims, "", " ")

			if err != nil {
				return "", err
			}

			return string(b), nil
		}
	}

	return "", errors.New("Unauthorized")

}
