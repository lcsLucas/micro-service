package pedidos

import (
	"context"

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

func (r *repository) Get(ctx context.Context, id int64) (Pedido, error) {
	p := Pedido{}
	var itens []Itens

	err := r.db.Debug().WithContext(ctx).Model(Pedido{}).Where("id=?", id).Take(&p).Error

	if err != nil {
		return Pedido{}, err
	}

	err = r.db.Debug().WithContext(ctx).Model(Itens{}).Where("idpedidos=?", id).Find(&itens).Error
	if err != nil {
		return Pedido{}, err
	}

	p.Itens = itens
	return p, nil
}
