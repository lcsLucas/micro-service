package migrations

import (
	"log"

	"github.com/lcslucas/micro-service/pedidos"
	"gorm.io/gorm"
)

func ExecMigrationPedidos(db *gorm.DB) error {
	var err error

	check := db.Migrator().HasTable(pedidos.Pedido{})
	if !check {
		err = db.Debug().Migrator().AutoMigrate(pedidos.Pedido{})
		if err != nil {
			return err
		} else {
			log.Println(`Tabela de "pedidos" gerada com sucesso.`)
		}
	}

	check = db.Migrator().HasTable(pedidos.Itens{})
	if !check {
		// s.DB.Debug().AutoMigrate(&models.Post{}).AddForeignKey("usuario_id", "users(id)", "CASCADE", "CASCADE") //database migration
		// db.Migrator().CreateConstraint(usuarios.Usuario{}, "ID")
		err = db.Debug().Migrator().AutoMigrate(pedidos.Itens{})
		if err != nil {
			return err
		} else {
			log.Println(`Tabela de "itens do pedido (pedido_itens)" gerada com sucesso.`)
		}
	}

	return nil
}
