package routers

import (
	"context"
	"encoding/json"
	"github.com/acossovich/twitterGo/bd"
	"github.com/acossovich/twitterGo/jwt"
	"github.com/acossovich/twitterGo/models"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"time"
)

func Login(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Usuario y/o contraseña invalidos" + err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "El email del usuario es requerido"
		return r
	}
	userData, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Usuario y/o contraseña invalidos"
		return r
	}

	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Ocurrio un error al intentar generar el Token correspondiente >" + err.Error()
		return r
	}

	respuesta := models.RespLogin{
		Token: jwtKey,
	}

	token, err := json.Marshal(respuesta)
	if err != nil {
		r.Message = "Ocurrio un error al intentar formatear el Token a JSON >" + err.Error()
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookieStr := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieStr,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
