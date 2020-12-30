package usuarios

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	clientMongo *mongo.Client
	logger      log.Logger
}

func NewRepository(client *mongo.Client, logger log.Logger) Repository {
	return &repository{
		clientMongo: client,
		logger:      log.With(logger, "repository", "sql"),
	}
}

func (r *repository) Get(ctx context.Context, id int) (Usuario, error) {
	db := r.clientMongo.Database("pedidos")
	collection := db.Collection("usuarios")

	filter := bson.M{"id": id}

	u := Usuario{}

	err := collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		return Usuario{}, err
	}

	fmt.Println(u)

	/*
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			return Usuario{}, err
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			u := Usuario{}
			err = cur.Decode(&u)
			if err != nil {
				return Usuario{}, err
			}

			fmt.Println(u)
		}
	*/

	return u, nil
}

/*
func (r *repository) Get(ctx context.Context, id int) (Usuario, error) {
	u := Usuario{}
	err := r.db.Debug().WithContext(ctx).Model(Usuario{}).Where("id=?", id).Take(&u).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Usuario{}, errors.New("Usuário não encontrado")
		}

		return Usuario{}, err
	}

	return u, nil
}
*/
