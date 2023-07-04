package jwt

import (
	"errors"
	"github.com/acossovich/twitterGo/models"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var Email, Usuario string

// Funcion para procesar el token
func ProcesoToken(token, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(token, "bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		// Rutina que chequea contra la BD
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token invalido")
	}

	return &claims, true, string(""), err
}
