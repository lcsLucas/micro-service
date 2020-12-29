package pedidos

import (
	"context"

	"github.com/lcslucas/micro-service/usuarios"
)

type Tabler interface {
	TableName() string
}

type Pedido struct {
	ID         int64            `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Data       string           `gorm:"size:10" json:"data"`
	ValorTotal float32          `json:"valor_total"`
	Pagamento  string           `gorm:"size:100" json:"pagamento,omitempty"`
	Usu        usuarios.Usuario `gorm:"ForeignKey:idusuarios;association_foreignkey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"usuario"`
	Idusuarios int              `json:"idusuarios"`
	Itens      []Itens          `gorm:" constraint: OnUpdate: CASCADE, OnDelete: CASCADE; " json:"itens"`
}

func (Pedido) TableName() string {
	return "pedidos"
}

type Itens struct {
	ID        int64   `gorm:"primary_key;auto_increment" json:"id"`
	Descricao string  `gorm:"size:255" json:"descricao"`
	Preco     float32 `json:"preco"`
	Qtde      uint32  `json:"qtde,omitempty"`
	PedidoID  int64
}

func (Itens) TableName() string {
	return "pedidos_itens"
}

type Repository interface {
	Get(ctx context.Context, id int64) (Pedido, error)
}
