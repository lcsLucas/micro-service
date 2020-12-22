package usuarios

import (
	"context"

	"github.com/go-kit/kit/log/level"

	"github.com/go-kit/kit/log"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func (s service) Get(ctx context.Context, id int) (Usuario, error) {
	logger := log.With(s.logger, "method", "Get")

	u, err := s.repository.Get(ctx, id)

	if err != nil {
		level.Error(logger).Log("[error]", err)
		return Usuario{}, err
	}

	level.Info(logger).Log("User[id]=", id)
	return u, nil
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}
