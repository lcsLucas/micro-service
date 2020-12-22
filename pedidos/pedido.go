package pedidos

import (
	"context"

	"github.com/lcslucas/micro-service/usuarios"
)

type Itens struct {
	ID        int64   `json:"id"`
	Descricao string  `json:"descricao"`
	Preco     float32 `json:"preco"`
	Qtde      uint32  `json:"qtde,omitempty"`
}

type Pedido struct {
	ID         int64            `json:"id, omitempty"`
	Data       string           `json:"data"`
	ValorTotal float32          `json:"valor_total"`
	Pagamento  string           `json:"pagamento,omitempty"`
	Usu        usuarios.Usuario `json:"usuario"`
	Itens      []Itens          `json:"itens"`
}

type Repository interface {
	Get(ctx context.Context, id int64) (Pedido, error)
}
