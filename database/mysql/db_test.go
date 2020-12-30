package mysql

import (
	"testing"

	"github.com/lcslucas/micro-service/config"
)

func TestDB(t *testing.T) {

	c := config.ConfigDB{
		Host:     "localhost",
		User:     "root",
		Password: "",
		Port:     "3306",
		DBName:   "pedidos",
	}

	_, err := Connect(c)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

/*

c.User, c.Password, c.Host, c.Port, c.DBName

USU_DB_HOST=localhost
USU_DB_DRIVER=mysql
USU_DB_USER=root
USU_DB_PASSWORD=""
USU_DB_NAME=pedidos
USU_DB_PORT=3306
*/
