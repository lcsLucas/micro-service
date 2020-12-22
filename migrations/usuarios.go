package migrations

import (
	"log"

	"github.com/lcslucas/micro-service/usuarios"
	"gorm.io/gorm"
)

func ExecMigrationUsuarios(db *gorm.DB) error {
	var err error

	check := db.Migrator().HasTable(usuarios.Usuario{})
	if !check {
		err = db.Debug().Migrator().AutoMigrate(usuarios.Usuario{})
		if err != nil {
			return err
		} else {
			log.Println(`Tabela de "usuarios" gerada com sucesso.`)
		}
	}

	return nil
}
