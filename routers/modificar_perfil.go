package routers

import (
	"context"
	"encoding/json"
	"github.com/acossovich/twitterGo/bd"
	"github.com/acossovich/twitterGo/models"
)

func ModificarPerfil(ctx context.Context, claims models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	var t models.Usuario

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Datos incorrectos" + err.Error()
		return r
	}

	status, err := bd.ModificoRegistro(t, claims.ID.Hex())
	if err != nil {
		r.Message = "Ocurrio un error al intentar modificar el registro. " + err.Error()
		return r
	}
	if !status {
		r.Message = "No se ha logrado modificar el registro del usuario. "
		return r
	}

	r.Status = 200
	r.Message = "Registro modificado correctamente"
	return r
}
