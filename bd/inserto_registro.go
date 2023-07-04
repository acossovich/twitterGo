package bd

import (
	"context"
	"github.com/acossovich/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(usuario models.Usuario) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := col.InsertOne(ctx, usuario)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
