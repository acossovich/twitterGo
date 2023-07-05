package routers

import (
	"context"
	"encoding/json"
	"github.com/acossovich/twitterGo/bd"
	"github.com/acossovich/twitterGo/models"
	"time"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.RespApi {
	var mensaje models.Tweet
	var r models.RespApi
	r.Status = 400
	idUsuario := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)
	if err != nil {
		r.Message = "Ocurrio un error al intentar decodificar el body " + err.Error()
		return r
	}

	registro := models.GraboTweet{
		UserID:  idUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		r.Message = "Ocurrio un error al intentar insertar el registro " + err.Error()
		return r
	}
	if !status {
		r.Message = "No se ha logrado insertar el registro del tweet "
		return r
	}

	r.Status = 200
	r.Message = "Tweet creado correctamente"
	return r
}
