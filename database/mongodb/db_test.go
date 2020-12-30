package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/lcslucas/micro-service/config"
	"github.com/lcslucas/micro-service/usuarios"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDB(t *testing.T) {

	c := config.ConfigDB{
		Host:     "localhost",
		User:     "root",
		Password: "",
		Port:     "27017",
		DBName:   "pedidos",
	}

	_, err := Connect(nil, c)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestInsert(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	c := config.ConfigDB{
		Host:     "localhost",
		User:     "root",
		Password: "",
		Port:     "27017",
		DBName:   "pedidos",
	}

	clientMongo, err := Connect(ctx, c)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer clientMongo.Disconnect(ctx)

	db := clientMongo.Database(c.DBName)

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
		t.Fatalf(err.Error())
	}
}

func TestGetAll(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	c := config.ConfigDB{
		Host:     "localhost",
		User:     "root",
		Password: "",
		Port:     "27017",
		DBName:   "pedidos",
	}

	clientMongo, err := Connect(ctx, c)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer clientMongo.Disconnect(ctx)

	db := clientMongo.Database(c.DBName)
	collection := db.Collection("usuarios")

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer cur.Close(ctx)

	t.Log(cur)
}
