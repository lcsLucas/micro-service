package usuarios

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db:     db,
		logger: log.With(logger, "repository", "sql"),
	}
}

func (r *repository) Get(ctx context.Context, id int) (Usuario, error) {
	var rid int
	var nome, email string

	err := r.db.QueryRow("SELECT * FROM usuarios WHERE id=? LIMIT 1", id).Scan(&rid, &nome, &email)
	if err != nil {
		return Usuario{}, err
	}

	return Usuario{
		ID:    rid,
		Nome:  nome,
		Email: email,
	}, nil
}
