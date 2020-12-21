package pedidos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/lcslucas/lojavirtual/usuarios"
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

func (r *repository) Get(ctx context.Context, id int64) (Pedido, error) {
	var rid, rusu int64
	var rdata, rpagto string
	var rvalor float32

	err := r.db.QueryRowContext(ctx, "SELECT id, data, idusuarios, valor_total, pagamento FROM pedidos WHERE id=? LIMIT 1", id).Scan(&rid, &rdata, &rusu, &rvalor, &rpagto)
	if err != nil {
		return Pedido{}, err
	}

	rows, err := r.db.QueryContext(ctx, "SELECT id, descricao, valor, qtde FROM pedidos_itens WHERE idpedidos=?", id)
	if err != nil {
		return Pedido{}, err
	}

	var itens []Itens

	for rows.Next() {
		var id int64
		var descricao string
		var preco float32
		var qtde uint32

		err = rows.Scan(&id, &descricao, &preco, &qtde)
		if err != nil {
			fmt.Println(err)
		}

		itens = append(itens, Itens{
			ID:        id,
			Descricao: descricao,
			Preco:     preco,
			Qtde:      qtde,
		})

	}

	return Pedido{
		ID:         rid,
		Data:       rdata,
		ValorTotal: rvalor,
		Pagamento:  rpagto,
		Usu: usuarios.Usuario{
			ID: int(rusu),
		},
		Itens: itens,
	}, nil

}
