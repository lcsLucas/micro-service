package usuarios

import "context"

type Service interface {
	Get(ctx context.Context, id int) (Usuario, error)
}
