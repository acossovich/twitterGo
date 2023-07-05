package bd

import (
	"context"
	"github.com/acossovich/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(user models.Usuario, ID string) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(user.Nombre) > 0 {
		registro["nombre"] = user.Nombre
	}
	if len(user.Apellidos) > 0 {
		registro["apellidos"] = user.Apellidos
	}
	registro["fechaNacimiento"] = user.FechaNacimiento
	if len(user.Avatar) > 0 {
		registro["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		registro["banner"] = user.Banner
	}
	if len(user.Biografia) > 0 {
		registro["biografia"] = user.Biografia
	}
	if len(user.Ubicacion) > 0 {
		registro["ubicacion"] = user.Ubicacion
	}
	if len(user.SitioWeb) > 0 {
		registro["sitioWeb"] = user.SitioWeb
	}

	updtStr := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtStr)
	if err != nil {
		return false, err
	}

	return true, nil
}
