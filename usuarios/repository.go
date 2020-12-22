package usuarios

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) Repository {
	return &repository{
		db:     db,
		logger: log.With(logger, "repository", "sql"),
	}
}

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
