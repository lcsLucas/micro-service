package migrations

import (
	"context"

	"github.com/lcslucas/micro-service/usuarios"
	"go.mongodb.org/mongo-driver/mongo"
)

func ExecMigrationUsuarios(clientMongo *mongo.Client) error {
	var err error

	db := clientMongo.Database("pedidos")

	users := []interface{}{
		usuarios.Usuario{
			ID:    1,
			Nome:  "Lucas S. Rosa",
			Email: "lucas.tarta@hotmail.com",
		},
		usuarios.Usuario{
			ID:    2,
			Nome:  "Maria da Silva",
			Email: "mariadasilva@email.com",
		},
	}

	collection := db.Collection("usuarios")
	_, err = collection.InsertMany(context.TODO(), users)
	if err != nil {
		return err
	}

	return nil
}
