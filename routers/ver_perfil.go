package routers

import (
	"encoding/json"
	"fmt"
	"github.com/acossovich/twitterGo/bd"
	"github.com/acossovich/twitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entre a ver perfil")

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio"
		return r
	}

	profile, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Ocurrió un error al intentar buscar el registro " + err.Error()
		return r
	}

	resp, err := json.Marshal(profile)
	if err != nil {
		r.Status = 500
		r.Message = "Ocurrió un error al formatear los datos de los usuarios como JSON " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(resp)

	return r
}
