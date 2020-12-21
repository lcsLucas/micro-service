package pedidos

import "context"

type Service interface {
	Get(ctx context.Context, id int64) (Pedido, error)
}
